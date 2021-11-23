/*
 * Copyright (c) 2019 Huawei Technologies Co., Ltd.
 * MeshAccelerating is licensed under the Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *     http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
 * PURPOSE.
 * See the Mulan PSL v2 for more details.
 * Author: LemmyHuang
 * Create: 2021-09-17
 */

#include "bpf_log.h"
#include "cluster.h"
#include "endpoint.h"
#include "tail_call.h"

static inline
int cluster_handle_circuit_breaker(cluster_t *cluster)
{
	// TODO
	return 0;
}

static inline
endpoint_t *loadbanance_round_robin(load_assignment_t *load_assignment)
{
	map_key_t *map_key = NULL;

	map_key = &load_assignment->map_key_of_least_endpoint;

	map_key->index++;
	map_key->index %= load_assignment->map_key_of_endpoint.index;

	return map_lookup_endpoint(map_key);
}

static inline
endpoint_t *loadbanance_least_request(load_assignment_t *load_assignment)
{
	unsigned i;
	map_key_t map_key;
	endpoint_t *endpoint = NULL;
	
	map_key = load_assignment->map_key_of_endpoint;

	for (i = 0; i < MAP_SIZE_OF_PER_ENDPOINT; i++) {
		map_key.index = i;

		endpoint_t *ep = map_lookup_endpoint(&map_key);
		if (ep == NULL)
			break;

		if ((endpoint == NULL) || (endpoint->lb_conn_num > ep->lb_conn_num))
			endpoint = ep;
	}

	// TODO: -1 when disconn
	if (endpoint != NULL)
		endpoint->lb_conn_num++;
	return endpoint;
}

static inline
int cluster_handle_loadbanance(ctx_buff_t *ctx, load_assignment_t *load_assignment)
{
	endpoint_t *endpoint = NULL;

	switch (load_assignment->lb_policy) {
		case LB_POLICY_LEAST_REQUEST:
			endpoint = loadbanance_least_request(load_assignment);
			break;
		case LB_POLICY_ROUND_ROBIN:
			endpoint = loadbanance_round_robin(load_assignment);
			break;
		case LB_POLICY_RANDOM:
			// TODO
			break;
		default:
			BPF_LOG(ERR, KMESH, "load_assignment lb_policy is wrong\n");
			break;
	}

	if (endpoint == NULL)
		return -EAGAIN;

	SET_CTX_ADDRESS(ctx, &endpoint->address);

	return 0;
}

SEC_TAIL(KMESH_SOCKET_CALLS, KMESH_TAIL_CALL_CLUSTER)
int cluster_manager(ctx_buff_t *ctx)
{
	int ret;
	map_key_t *pkey = NULL;
	cluster_t *cluster = NULL;

	DECLARE_VAR_ADDRESS(ctx, address);

	pkey = kmesh_tail_lookup_ctx(&address);
	if (pkey == NULL)
		return convert_sock_errno(ENOENT);

	cluster = map_lookup_cluster(pkey);
	kmesh_tail_delete_ctx(&address);
	if (cluster == NULL)
		return convert_sock_errno(ENOENT);

	if (cluster_handle_circuit_breaker(cluster) != 0)
		return convert_sock_errno(EBUSY);

	ret = cluster_handle_loadbanance(ctx, &cluster->load_assignment);
	return convert_sock_errno(ret);
}

char _license[] SEC("license") = "GPL";
int _version SEC("version") = 1;