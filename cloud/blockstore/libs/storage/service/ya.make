LIBRARY()

SRCS(
    service.cpp
    service_actor.cpp
    service_actor_actions_backup_disk_registry_state.cpp
    service_actor_actions_change_disk_device.cpp
    service_actor_actions_change_storage_config.cpp
    service_actor_actions_check_blob.cpp
    service_actor_actions_cms.cpp
    service_actor_actions_compact_range.cpp
    service_actor_actions_configure_volume_balancer.cpp
    service_actor_actions_create_disk_from_devices.cpp
    service_actor_actions_delete_checkpoint_data.cpp
    service_actor_actions_describe_blocks.cpp
    service_actor_actions_describe.cpp
    service_actor_actions_disk_registry_change_state.cpp
    service_actor_actions_drain_node.cpp
    service_actor_actions_finish_fill_disk.cpp
    service_actor_actions_flush_profile_log.cpp
    service_actor_actions_get_dependent_disks.cpp
    service_actor_actions_get_compaction_status.cpp
    service_actor_actions_get_disk_agent_node_id.cpp
    service_actor_actions_get_diskregistry_tablet_info.cpp
    service_actor_actions_get_nameservice_config.cpp
    service_actor_actions_get_partition_info.cpp
    service_actor_actions_kill_tablet.cpp
    service_actor_actions_migration_disk_registry_device.cpp
    service_actor_actions_modify_tags.cpp
    service_actor_actions_reallocate_disk.cpp
    service_actor_actions_reassign_disk_registry.cpp
    service_actor_actions_rebase_volume.cpp
    service_actor_actions_rebind_volumes.cpp
    service_actor_actions_rebuild_metadata.cpp
    service_actor_actions_replace_device.cpp
    service_actor_actions_reset_tablet.cpp
    service_actor_actions_restore_disk_registry_state.cpp
    service_actor_actions_scan_disk.cpp
    service_actor_actions_set_user_id.cpp
    service_actor_actions_setup_channels.cpp
    service_actor_actions_suspend_device.cpp
    service_actor_actions_suspend_disk_agent.cpp
    service_actor_actions_update_disk_block_size.cpp
    service_actor_actions_update_disk_replica_count.cpp
    service_actor_actions_update_disk_registry_params.cpp
    service_actor_actions_update_placement_group_settings.cpp
    service_actor_actions_update_used_blocks.cpp
    service_actor_actions_update_volume_params.cpp
    service_actor_actions_wait_dependent_disks_switched_node.cpp
    service_actor_actions_writable_state.cpp
    service_actor_actions.cpp
    service_actor_alter.cpp
    service_actor_assign.cpp
    service_actor_balancer_stats.cpp
    service_actor_check_liveness.cpp
    service_actor_client_stats.cpp
    service_actor_create.cpp
    service_actor_create_from_device.cpp
    service_actor_describe.cpp
    service_actor_describe_disk_registry_config.cpp
    service_actor_describe_model.cpp
    service_actor_destroy.cpp
    service_actor_forward.cpp
    service_actor_forward_to_disk_registry.cpp
    service_actor_list.cpp
    service_actor_monitoring.cpp
    service_actor_monitoring_binding.cpp
    service_actor_monitoring_clients.cpp
    service_actor_monitoring_search.cpp
    service_actor_monitoring_unmount.cpp
    service_actor_mount.cpp
    service_actor_readblocks.cpp
    service_actor_sync_manually_preempted_volumes.cpp
    service_actor_unmount.cpp
    service_actor_update_disk_registry_config.cpp
    service_actor_volume_binding.cpp
    service_actor_writeblocks.cpp
    service_counters.cpp
    service_state.cpp
    volume_client_actor.cpp
    volume_session_actor.cpp
    volume_session_actor_binding.cpp
    volume_session_actor_mount.cpp
    volume_session_actor_start.cpp
    volume_session_actor_stop.cpp
    volume_session_actor_unmount.cpp
    volume_timed_delivery.cpp
)

PEERDIR(
    cloud/blockstore/libs/common
    cloud/blockstore/libs/diagnostics
    cloud/blockstore/libs/discovery
    cloud/blockstore/libs/encryption
    cloud/blockstore/libs/kikimr
    cloud/blockstore/libs/storage/api
    cloud/blockstore/libs/storage/core
    cloud/blockstore/libs/storage/protos
    cloud/blockstore/libs/storage/service/model
    cloud/blockstore/libs/storage/volume
    cloud/blockstore/private/api/protos

    cloud/storage/core/libs/api
    cloud/storage/core/libs/common

    contrib/ydb/library/actors/core
    library/cpp/json
    library/cpp/monlib/service/pages

    contrib/ydb/core/base
    contrib/ydb/core/mind
    contrib/ydb/core/mon
    contrib/ydb/core/node_whiteboard
    contrib/ydb/core/protos
    contrib/ydb/core/tablet
    contrib/ydb/core/tx/schemeshard

    contrib/libs/protobuf
)

END()

RECURSE(
    model
)

RECURSE_FOR_TESTS(
    ut
)
