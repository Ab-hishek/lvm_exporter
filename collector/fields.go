package collector

type FieldType string

const (
	FieldTypeSTR      FieldType = "STR"
	FieldTypeNUM      FieldType = "NUM"
	FieldTypeBIN      FieldType = "BIN"
	FieldTypeSIZ      FieldType = "SIZ"
	FieldTypePCT      FieldType = "PCT"
	FieldTypeTIM      FieldType = "TIM"
	FieldTypeSNUM     FieldType = "SNUM"
	FieldTypeSTR_LIST FieldType = "STR_LIST"
)

type Field struct {
	Type        FieldType
	Name        string
	Description string
}

var lvmFields = map[string]Field{
	"lv_uuid": Field{
		Type:        FieldTypeSTR,
		Name:        "LV UUID",
		Description: "Unique identifier.",
	},
	"lv_name": Field{
		Type:        FieldTypeSTR,
		Name:        "LV",
		Description: "Name.  LVs created for internal use are enclosed in brackets.",
	},
	"lv_full_name": Field{
		Type:        FieldTypeSTR,
		Name:        "LV",
		Description: "Full name of LV including its VG, namely VG/LV.",
	},
	"lv_path": Field{
		Type:        FieldTypeSTR,
		Name:        "Path",
		Description: "Full pathname for LV. Blank for internal LVs.",
	},
	"lv_dm_path": Field{
		Type:        FieldTypeSTR,
		Name:        "DMPath",
		Description: "Internal device-mapper pathname for LV (in /dev/mapper directory).",
	},
	"lv_parent": Field{
		Type:        FieldTypeSTR,
		Name:        "Parent",
		Description: "For LVs that are components of another LV, the parent LV.",
	},
	"lv_layout": Field{
		Type:        FieldTypeSTR_LIST,
		Name:        "Layout",
		Description: "LV layout.",
	},
	"lv_role": Field{
		Type:        FieldTypeSTR_LIST,
		Name:        "Role",
		Description: "LV role.",
	},
	"lv_initial_image_sync": Field{
		Type:        FieldTypeBIN,
		Name:        "InitImgSync",
		Description: "Set if mirror/RAID images underwent initial resynchronization.",
	},
	"lv_image_synced": Field{
		Type:        FieldTypeBIN,
		Name:        "ImgSynced",
		Description: "Set if mirror/RAID image is synchronized.",
	},
	"lv_merging": Field{
		Type:        FieldTypeBIN,
		Name:        "Merging",
		Description: "Set if snapshot LV is being merged to origin.",
	},
	"lv_converting": Field{
		Type:        FieldTypeBIN,
		Name:        "Converting",
		Description: "Set if LV is being converted.",
	},
	"lv_allocation_policy": Field{
		Type:        FieldTypeSTR,
		Name:        "AllocPol",
		Description: "LV allocation policy.",
	},
	"lv_allocation_locked": Field{
		Type:        FieldTypeBIN,
		Name:        "AllocLock",
		Description: "Set if LV is locked against allocation changes.",
	},
	"lv_fixed_minor": Field{
		Type:        FieldTypeBIN,
		Name:        "FixMin",
		Description: "Set if LV has fixed minor number assigned.",
	},
	"lv_skip_activation": Field{
		Type:        FieldTypeBIN,
		Name:        "SkipAct",
		Description: "Set if LV is skipped on activation.",
	},
	"lv_when_full": Field{
		Type:        FieldTypeSTR,
		Name:        "WhenFull",
		Description: "For thin pools, behavior when full.",
	},
	"lv_active": Field{
		Type:        FieldTypeSTR,
		Name:        "Active",
		Description: "Active state of the LV.",
	},
	"lv_active_locally": Field{
		Type:        FieldTypeBIN,
		Name:        "ActLocal",
		Description: "Set if the LV is active locally.",
	},
	"lv_active_remotely": Field{
		Type:        FieldTypeBIN,
		Name:        "ActRemote",
		Description: "Set if the LV is active remotely.",
	},
	"lv_active_exclusively": Field{
		Type:        FieldTypeBIN,
		Name:        "ActExcl",
		Description: "Set if the LV is active exclusively.",
	},
	"lv_major": Field{
		Type:        FieldTypeSNUM,
		Name:        "Maj",
		Description: "Persistent major number or -1 if not persistent.",
	},
	"lv_minor": Field{
		Type:        FieldTypeSNUM,
		Name:        "Min",
		Description: "Persistent minor number or -1 if not persistent.",
	},
	"lv_read_ahead": Field{
		Type:        FieldTypeSIZ,
		Name:        "Rahead",
		Description: "Read ahead setting in current units.",
	},
	"lv_size": Field{
		Type:        FieldTypeSIZ,
		Name:        "LSize",
		Description: "Size of LV in current units.",
	},
	"lv_metadata_size": Field{
		Type:        FieldTypeSIZ,
		Name:        "MSize",
		Description: "For thin and cache pools, the size of the LV that holds the metadata.",
	},
	"seg_count": Field{
		Type:        FieldTypeNUM,
		Name:        "#Seg",
		Description: "Number of segments in LV.",
	},
	"origin": Field{
		Type:        FieldTypeSTR,
		Name:        "Origin",
		Description: "For snapshots and thins, the origin device of this LV.",
	},
	"origin_uuid": Field{
		Type:        FieldTypeSTR,
		Name:        "Origin UUID",
		Description: "For snapshots and thins, the UUID of origin device of this LV.",
	},
	"origin_size": Field{
		Type:        FieldTypeSIZ,
		Name:        "OSize",
		Description: "For snapshots, the size of the origin device of this LV.",
	},
	"lv_ancestors": Field{
		Type:        FieldTypeSTR_LIST,
		Name:        "Ancestors",
		Description: "LV ancestors ignoring any stored history of the ancestry chain.",
	},
	"lv_full_ancestors": Field{
		Type:        FieldTypeSTR_LIST,
		Name:        "FAncestors",
		Description: "LV ancestors including stored history of the ancestry chain.",
	},
	"lv_descendants": Field{
		Type:        FieldTypeSTR_LIST,
		Name:        "Descendants",
		Description: "LV descendants ignoring any stored history of the ancestry chain.",
	},
	"lv_full_descendants": Field{
		Type:        FieldTypeSTR_LIST,
		Name:        "FDescendants",
		Description: "LV descendants including stored history of the ancestry chain.",
	},
	"raid_mismatch_count": Field{
		Type:        FieldTypeNUM,
		Name:        "Mismatches",
		Description: "For RAID, number of mismatches found or repaired.",
	},
	"raid_sync_action": Field{
		Type:        FieldTypeSTR,
		Name:        "SyncAction",
		Description: "For RAID, the current synchronization action being performed.",
	},
	"raid_write_behind": Field{
		Type:        FieldTypeNUM,
		Name:        "WBehind",
		Description: "For RAID1, the number of outstanding writes allowed to writemostly devices.",
	},
	"raid_min_recovery_rate": Field{
		Type:        FieldTypeNUM,
		Name:        "MinSync",
		Description: "For RAID1, the minimum recovery I/O load in kiB/sec/disk.",
	},
	"raid_max_recovery_rate": Field{
		Type:        FieldTypeNUM,
		Name:        "MaxSync",
		Description: "For RAID1, the maximum recovery I/O load in kiB/sec/disk.",
	},
	"raidintegritymode": Field{
		Type:        FieldTypeSTR,
		Name:        "IntegMode",
		Description: "The integrity mode",
	},
	"raidintegrityblocksize": Field{
		Type:        FieldTypeNUM,
		Name:        "IntegBlkSize",
		Description: "The integrity block size",
	},
	"integritymismatches": Field{
		Type:        FieldTypeNUM,
		Name:        "IntegMismatches",
		Description: "The number of integrity mismatches.",
	},
	"move_pv": Field{
		Type:        FieldTypeSTR,
		Name:        "Move",
		Description: "For pvmove, Source PV of temporary LV created by pvmove.",
	},
	"move_pv_uuid": Field{
		Type:        FieldTypeSTR,
		Name:        "Move UUID",
		Description: "For pvmove, the UUID of Source PV of temporary LV created by pvmove.",
	},
	"convert_lv": Field{
		Type:        FieldTypeSTR,
		Name:        "Convert",
		Description: "For lvconvert, Name of temporary LV created by lvconvert.",
	},
	"convert_lv_uuid": Field{
		Type:        FieldTypeSTR,
		Name:        "Convert UUID",
		Description: "For lvconvert, UUID of temporary LV created by lvconvert.",
	},
	"mirror_log": Field{
		Type:        FieldTypeSTR,
		Name:        "Log",
		Description: "For mirrors, the LV holding the synchronisation log.",
	},
	"mirror_log_uuid": Field{
		Type:        FieldTypeSTR,
		Name:        "Log UUID",
		Description: "For mirrors, the UUID of the LV holding the synchronisation log.",
	},
	"data_lv": Field{
		Type:        FieldTypeSTR,
		Name:        "Data",
		Description: "For cache/thin/vdo pools, the LV holding the associated data.",
	},
	"data_lv_uuid": Field{
		Type:        FieldTypeSTR,
		Name:        "Data UUID",
		Description: "For cache/thin/vdo pools, the UUID of the LV holding the associated data.",
	},
	"metadata_lv": Field{
		Type:        FieldTypeSTR,
		Name:        "Meta",
		Description: "For cache/thin pools, the LV holding the associated metadata.",
	},
	"metadata_lv_uuid": Field{
		Type:        FieldTypeSTR,
		Name:        "Meta UUID",
		Description: "For cache/thin pools, the UUID of the LV holding the associated metadata.",
	},
	"pool_lv": Field{
		Type:        FieldTypeSTR,
		Name:        "Pool",
		Description: "For cache/thin/vdo volumes, the cache/thin/vdo pool LV for this volume.",
	},
	"pool_lv_uuid": Field{
		Type:        FieldTypeSTR,
		Name:        "Pool UUID",
		Description: "For cache/thin/vdo volumes, the UUID of the cache/thin/vdo pool LV for this volume.",
	},
	"lv_tags": Field{
		Type:        FieldTypeSTR_LIST,
		Name:        "LV Tags",
		Description: "Tags, if any.",
	},
	"lv_profile": Field{
		Type:        FieldTypeSTR,
		Name:        "LProfile",
		Description: "Configuration profile attached to this LV.",
	},
	"lv_lockargs": Field{
		Type:        FieldTypeSTR,
		Name:        "LLockArgs",
		Description: "Lock args of the LV used by lvmlockd.",
	},
	"lv_time": Field{
		Type:        FieldTypeTIM,
		Name:        "CTime",
		Description: "Creation time of the LV, if known",
	},
	"lv_time_removed": Field{
		Type:        FieldTypeTIM,
		Name:        "RTime",
		Description: "Removal time of the LV, if known",
	},
	"lv_host": Field{
		Type:        FieldTypeSTR,
		Name:        "Host",
		Description: "Creation host of the LV, if known.",
	},
	"lv_modules": Field{
		Type:        FieldTypeSTR_LIST,
		Name:        "Modules",
		Description: "Kernel device-mapper modules required for this LV.",
	},
	"lv_historical": Field{
		Type:        FieldTypeBIN,
		Name:        "Historical",
		Description: "Set if the LV is historical.",
	},
	"lv_kernel_major": Field{
		Type:        FieldTypeSNUM,
		Name:        "KMaj",
		Description: "Currently assigned major number or -1 if LV is not active.",
	},
	"lv_kernel_minor": Field{
		Type:        FieldTypeSNUM,
		Name:        "KMin",
		Description: "Currently assigned minor number or -1 if LV is not active.",
	},
	"lv_kernel_read_ahead": Field{
		Type:        FieldTypeSIZ,
		Name:        "KRahead",
		Description: "Currently-in-use read ahead setting in current units.",
	},
	"lv_permissions": Field{
		Type:        FieldTypeSTR,
		Name:        "LPerms",
		Description: "LV permissions.",
	},
	"lv_suspended": Field{
		Type:        FieldTypeBIN,
		Name:        "Suspended",
		Description: "Set if LV is suspended.",
	},
	"lv_live_table": Field{
		Type:        FieldTypeBIN,
		Name:        "LiveTable",
		Description: "Set if LV has live table present.",
	},
	"lv_inactive_table": Field{
		Type:        FieldTypeBIN,
		Name:        "InactiveTable",
		Description: "Set if LV has inactive table present.",
	},
	"lv_device_open": Field{
		Type:        FieldTypeBIN,
		Name:        "DevOpen",
		Description: "Set if LV device is open.",
	},
	"data_percent": Field{
		Type:        FieldTypePCT,
		Name:        "Data%",
		Description: "For snapshot, cache and thin pools and volumes, the percentage full if LV is active.",
	},
	"snap_percent": Field{
		Type:        FieldTypePCT,
		Name:        "Snap%",
		Description: "For snapshots, the percentage full if LV is active.",
	},
	"metadata_percent": Field{
		Type:        FieldTypePCT,
		Name:        "Meta%",
		Description: "For cache and thin pools, the percentage of metadata full if LV is active.",
	},
	"copy_percent": Field{
		Type:        FieldTypePCT,
		Name:        "Cpy%Sync",
		Description: "For Cache, RAID, mirrors and pvmove, current percentage in-sync.",
	},
	"sync_percent": Field{
		Type:        FieldTypePCT,
		Name:        "Cpy%Sync",
		Description: "For Cache, RAID, mirrors and pvmove, current percentage in-sync.",
	},
	"cache_total_blocks": Field{
		Type:        FieldTypeNUM,
		Name:        "CacheTotalBlocks",
		Description: "Total cache blocks.",
	},
	"cache_used_blocks": Field{
		Type:        FieldTypeNUM,
		Name:        "CacheUsedBlocks",
		Description: "Used cache blocks.",
	},
	"cache_dirty_blocks": Field{
		Type:        FieldTypeNUM,
		Name:        "CacheDirtyBlocks",
		Description: "Dirty cache blocks.",
	},
	"cache_read_hits": Field{
		Type:        FieldTypeNUM,
		Name:        "CacheReadHits",
		Description: "Cache read hits.",
	},
	"cache_read_misses": Field{
		Type:        FieldTypeNUM,
		Name:        "CacheReadMisses",
		Description: "Cache read misses.",
	},
	"cache_write_hits": Field{
		Type:        FieldTypeNUM,
		Name:        "CacheWriteHits",
		Description: "Cache write hits.",
	},
	"cache_write_misses": Field{
		Type:        FieldTypeNUM,
		Name:        "CacheWriteMisses",
		Description: "Cache write misses.",
	},
	"kernel_cache_settings": Field{
		Type:        FieldTypeSTR_LIST,
		Name:        "KCacheSettings",
		Description: "Cache settings/parameters as set in kernel, including default values (cached segments only).",
	},
	"kernel_cache_policy": Field{
		Type:        FieldTypeSTR,
		Name:        "KCachePolicy",
		Description: "Cache policy used in kernel.",
	},
	"kernel_metadata_format": Field{
		Type:        FieldTypeNUM,
		Name:        "KMFmt",
		Description: "Cache metadata format used in kernel.",
	},
	"lv_health_status": Field{
		Type:        FieldTypeSTR,
		Name:        "Health",
		Description: "LV health status.",
	},
	"kernel_discards": Field{
		Type:        FieldTypeSTR,
		Name:        "KDiscards",
		Description: "For thin pools, how discards are handled in kernel.",
	},
	"lv_check_needed": Field{
		Type:        FieldTypeBIN,
		Name:        "CheckNeeded",
		Description: "For thin pools and cache volumes, whether metadata check is needed.",
	},
	"lv_merge_failed": Field{
		Type:        FieldTypeBIN,
		Name:        "MergeFailed",
		Description: "Set if snapshot merge failed.",
	},
	"lv_snapshot_invalid": Field{
		Type:        FieldTypeBIN,
		Name:        "SnapInvalid",
		Description: "Set if snapshot LV is invalid.",
	},
	"vdo_operating_mode": Field{
		Type:        FieldTypeSTR,
		Name:        "VDOOperatingMode",
		Description: "For vdo pools, its current operating mode.",
	},
	"vdo_compression_state": Field{
		Type:        FieldTypeSTR,
		Name:        "VDOCompressionState",
		Description: "For vdo pools, whether compression is running.",
	},
	"vdo_index_state": Field{
		Type:        FieldTypeSTR,
		Name:        "VDOIndexState",
		Description: "For vdo pools, state of index for deduplication.",
	},
	"vdo_used_size": Field{
		Type:        FieldTypeNUM,
		Name:        "VDOUsedSize",
		Description: "For vdo pools, currently used space.",
	},
	"vdo_saving_percent": Field{
		Type:        FieldTypeNUM,
		Name:        "VDOSaving%",
		Description: "For vdo pools, percentage of saved space.",
	},
	"writecache_total_blocks": Field{
		Type:        FieldTypeNUM,
		Name:        "WCacheTotalBlocks",
		Description: "Total writecache blocks.",
	},
	"writecache_free_blocks": Field{
		Type:        FieldTypeNUM,
		Name:        "WCacheFreeBlocks",
		Description: "Total writecache free blocks.",
	},
	"writecache_writeback_blocks": Field{
		Type:        FieldTypeNUM,
		Name:        "WCacheWritebackBlocks",
		Description: "Total writecache writeback blocks.",
	},
	"writecache_error": Field{
		Type:        FieldTypeNUM,
		Name:        "WCacheErrors",
		Description: "Total writecache errors.",
	},
	"lv_attr": Field{
		Type:        FieldTypeSTR,
		Name:        "Attr",
		Description: "Various attributes - see man page.",
	},
	"pv_fmt": Field{
		Type:        FieldTypeSTR,
		Name:        "Fmt",
		Description: "Type of metadata.",
	},
	"pv_uuid": Field{
		Type:        FieldTypeSTR,
		Name:        "PV UUID",
		Description: "Unique identifier.",
	},
	"dev_size": Field{
		Type:        FieldTypeSIZ,
		Name:        "DevSize",
		Description: "Size of underlying device in current units.",
	},
	"pv_name": Field{
		Type:        FieldTypeSTR,
		Name:        "PV",
		Description: "Name.",
	},
	"pv_major": Field{
		Type:        FieldTypeSTR,
		Name:        "Maj",
		Description: "Device major number.",
	},
	"pv_minor": Field{
		Type:        FieldTypeSTR,
		Name:        "Min",
		Description: "Device minor number.",
	},
	"pv_mda_free": Field{
		Type:        FieldTypeSIZ,
		Name:        "PMdaFree",
		Description: "Free metadata area space on this device in current units.",
	},
	"pv_mda_size": Field{
		Type:        FieldTypeSIZ,
		Name:        "PMdaSize",
		Description: "Size of smallest metadata area on this device in current units.",
	},
	"pv_ext_vsn": Field{
		Type:        FieldTypeNUM,
		Name:        "PExtVsn",
		Description: "PV header extension version.",
	},
	"pe_start": Field{
		Type:        FieldTypeNUM,
		Name:        "1st PE",
		Description: "Offset to the start of data on the underlying device.",
	},
	"pv_size": Field{
		Type:        FieldTypeSIZ,
		Name:        "PSize",
		Description: "Size of PV in current units.",
	},
	"pv_free": Field{
		Type:        FieldTypeSIZ,
		Name:        "PFree",
		Description: "Total amount of unallocated space in current units.",
	},
	"pv_used": Field{
		Type:        FieldTypeSIZ,
		Name:        "Used",
		Description: "Total amount of allocated space in current units.",
	},
	"pv_attr": Field{
		Type:        FieldTypeSTR,
		Name:        "Attr",
		Description: "Various attributes - see man page.",
	},
	"pv_allocatable": Field{
		Type:        FieldTypeBIN,
		Name:        "Allocatable",
		Description: "Set if this device can be used for allocation.",
	},
	"pv_exported": Field{
		Type:        FieldTypeBIN,
		Name:        "Exported",
		Description: "Set if this device is exported.",
	},
	"pv_missing": Field{
		Type:        FieldTypeBIN,
		Name:        "Missing",
		Description: "Set if this device is missing in system.",
	},
	"pv_pe_count": Field{
		Type:        FieldTypeNUM,
		Name:        "PE",
		Description: "Total number of Physical Extents.",
	},
	"pv_pe_alloc_count": Field{
		Type:        FieldTypeNUM,
		Name:        "Alloc",
		Description: "Total number of allocated Physical Extents.",
	},
	"pv_tags": Field{
		Type:        FieldTypeSTR_LIST,
		Name:        "PV Tags",
		Description: "Tags, if any.",
	},
	"pv_mda_count": Field{
		Type:        FieldTypeNUM,
		Name:        "#PMda",
		Description: "Number of metadata areas on this device.",
	},
	"pv_mda_used_count": Field{
		Type:        FieldTypeNUM,
		Name:        "#PMdaUse",
		Description: "Number of metadata areas in use on this device.",
	},
	"pv_ba_start": Field{
		Type:        FieldTypeSIZ,
		Name:        "BA Start",
		Description: "Offset to the start of PV Bootloader Area on the underlying device in current units.",
	},
	"pv_ba_size": Field{
		Type:        FieldTypeSIZ,
		Name:        "BA Size",
		Description: "Size of PV Bootloader Area in current units.",
	},
	"pv_in_use": Field{
		Type:        FieldTypeBIN,
		Name:        "PInUse",
		Description: "Set if PV is used.",
	},
	"pv_duplicate": Field{
		Type:        FieldTypeBIN,
		Name:        "Duplicate",
		Description: "Set if PV is an unchosen duplicate.",
	},
	"vg_fmt": Field{
		Type:        FieldTypeSTR,
		Name:        "Fmt",
		Description: "Type of metadata.",
	},
	"vg_uuid": Field{
		Type:        FieldTypeSTR,
		Name:        "VG UUID",
		Description: "Unique identifier.",
	},
	"vg_name": Field{
		Type:        FieldTypeSTR,
		Name:        "VG",
		Description: "Name.",
	},
	"vg_attr": Field{
		Type:        FieldTypeSTR,
		Name:        "Attr",
		Description: "Various attributes - see man page.",
	},
	"vg_permissions": Field{
		Type:        FieldTypeSTR,
		Name:        "VPerms",
		Description: "VG permissions.",
	},
	"vg_extendable": Field{
		Type:        FieldTypeBIN,
		Name:        "Extendable",
		Description: "Set if VG is extendable.",
	},
	"vg_exported": Field{
		Type:        FieldTypeBIN,
		Name:        "Exported",
		Description: "Set if VG is exported.",
	},
	"vg_partial": Field{
		Type:        FieldTypeBIN,
		Name:        "Partial",
		Description: "Set if VG is partial.",
	},
	"vg_allocation_policy": Field{
		Type:        FieldTypeSTR,
		Name:        "AllocPol",
		Description: "VG allocation policy.",
	},
	"vg_clustered": Field{
		Type:        FieldTypeBIN,
		Name:        "Clustered",
		Description: "Set if VG is clustered.",
	},
	"vg_shared": Field{
		Type:        FieldTypeBIN,
		Name:        "Shared",
		Description: "Set if VG is shared.",
	},
	"vg_size": Field{
		Type:        FieldTypeSIZ,
		Name:        "VSize",
		Description: "Total size of VG in current units.",
	},
	"vg_free": Field{
		Type:        FieldTypeSIZ,
		Name:        "VFree",
		Description: "Total amount of free space in current units.",
	},
	"vg_sysid": Field{
		Type:        FieldTypeSTR,
		Name:        "SYS ID",
		Description: "System ID of the VG indicating which host owns it.",
	},
	"vg_systemid": Field{
		Type:        FieldTypeSTR,
		Name:        "System ID",
		Description: "System ID of the VG indicating which host owns it.",
	},
	"vg_lock_type": Field{
		Type:        FieldTypeSTR,
		Name:        "LockType",
		Description: "Lock type of the VG used by lvmlockd.",
	},
	"vg_lock_args": Field{
		Type:        FieldTypeSTR,
		Name:        "VLockArgs",
		Description: "Lock args of the VG used by lvmlockd.",
	},
	"vg_extent_size": Field{
		Type:        FieldTypeSIZ,
		Name:        "Ext",
		Description: "Size of Physical Extents in current units.",
	},
	"vg_extent_count": Field{
		Type:        FieldTypeNUM,
		Name:        "#Ext",
		Description: "Total number of Physical Extents.",
	},
	"vg_free_count": Field{
		Type:        FieldTypeNUM,
		Name:        "Free",
		Description: "Total number of unallocated Physical Extents.",
	},
	"max_lv": Field{
		Type:        FieldTypeNUM,
		Name:        "MaxLV",
		Description: "Maximum number of LVs allowed in VG or 0 if unlimited.",
	},
	"max_pv": Field{
		Type:        FieldTypeNUM,
		Name:        "MaxPV",
		Description: "Maximum number of PVs allowed in VG or 0 if unlimited.",
	},
	"pv_count": Field{
		Type:        FieldTypeNUM,
		Name:        "#PV",
		Description: "Number of PVs in VG.",
	},
	"vg_missing_pv_count": Field{
		Type:        FieldTypeNUM,
		Name:        "#PV Missing",
		Description: "Number of PVs in VG which are missing.",
	},
	"lv_count": Field{
		Type:        FieldTypeNUM,
		Name:        "#LV",
		Description: "Number of LVs.",
	},
	"snap_count": Field{
		Type:        FieldTypeNUM,
		Name:        "#SN",
		Description: "Number of snapshots.",
	},
	"vg_seqno": Field{
		Type:        FieldTypeNUM,
		Name:        "Seq",
		Description: "Revision number of internal metadata.  Incremented whenever it changes.",
	},
	"vg_tags": Field{
		Type:        FieldTypeSTR_LIST,
		Name:        "VG Tags",
		Description: "Tags, if any.",
	},
	"vg_profile": Field{
		Type:        FieldTypeSTR,
		Name:        "VProfile",
		Description: "Configuration profile attached to this VG.",
	},
	"vg_mda_count": Field{
		Type:        FieldTypeNUM,
		Name:        "#VMda",
		Description: "Number of metadata areas on this VG.",
	},
	"vg_mda_used_count": Field{
		Type:        FieldTypeNUM,
		Name:        "#VMdaUse",
		Description: "Number of metadata areas in use on this VG.",
	},
	"vg_mda_free": Field{
		Type:        FieldTypeSIZ,
		Name:        "VMdaFree",
		Description: "Free metadata area space for this VG in current units.",
	},
	"vg_mda_size": Field{
		Type:        FieldTypeSIZ,
		Name:        "VMdaSize",
		Description: "Size of smallest metadata area for this VG in current units.",
	},
	"vg_mda_copies": Field{
		Type:        FieldTypeNUM,
		Name:        "#VMdaCps",
		Description: "Target number of in use metadata areas in the VG.",
	},
	"segtype": Field{
		Type:        FieldTypeSTR,
		Name:        "Type",
		Description: "Type of LV segment.",
	},
	"stripes": Field{
		Type:        FieldTypeNUM,
		Name:        "#Str",
		Description: "Number of stripes or mirror/raid1 legs.",
	},
	"data_stripes": Field{
		Type:        FieldTypeNUM,
		Name:        "#DStr",
		Description: "Number of data stripes or mirror/raid1 legs.",
	},
	"reshape_len": Field{
		Type:        FieldTypeSIZ,
		Name:        "RSize",
		Description: "Size of out-of-place reshape space in current units.",
	},
	"reshape_len_le": Field{
		Type:        FieldTypeNUM,
		Name:        "RSize",
		Description: "Size of out-of-place reshape space in logical extents.",
	},
	"data_copies": Field{
		Type:        FieldTypeNUM,
		Name:        "#Cpy",
		Description: "Number of data copies.",
	},
	"data_offset": Field{
		Type:        FieldTypeNUM,
		Name:        "DOff",
		Description: "Data offset on each image device.",
	},
	"new_data_offset": Field{
		Type:        FieldTypeNUM,
		Name:        "NOff",
		Description: "New data offset after any reshape on each image device.",
	},
	"parity_chunks": Field{
		Type:        FieldTypeNUM,
		Name:        "#Par",
		Description: "Number of (rotating) parity chunks.",
	},
	"stripe_size": Field{
		Type:        FieldTypeSIZ,
		Name:        "Stripe",
		Description: "For stripes, amount of data placed on one device before switching to the next.",
	},
	"region_size": Field{
		Type:        FieldTypeSIZ,
		Name:        "Region",
		Description: "For mirrors/raids, the unit of data per leg when synchronizing devices.",
	},
	"chunk_size": Field{
		Type:        FieldTypeSIZ,
		Name:        "Chunk",
		Description: "For snapshots, the unit of data used when tracking changes.",
	},
	"thin_count": Field{
		Type:        FieldTypeNUM,
		Name:        "#Thins",
		Description: "For thin pools, the number of thin volumes in this pool.",
	},
	"discards": Field{
		Type:        FieldTypeSTR,
		Name:        "Discards",
		Description: "For thin pools, how discards are handled.",
	},
	"cache_metadata_format": Field{
		Type:        FieldTypeNUM,
		Name:        "CMFmt",
		Description: "For cache, metadata format in use.",
	},
	"cache_mode": Field{
		Type:        FieldTypeSTR,
		Name:        "CacheMode",
		Description: "For cache, how writes are cached.",
	},
	"zero": Field{
		Type:        FieldTypeBIN,
		Name:        "Zero",
		Description: "For thin pools and volumes, if zeroing is enabled.",
	},
	"transaction_id": Field{
		Type:        FieldTypeNUM,
		Name:        "TransId",
		Description: "For thin pools, the transaction id and creation transaction id for thins.",
	},
	"thin_id": Field{
		Type:        FieldTypeNUM,
		Name:        "ThId",
		Description: "For thin volume, the thin device id.",
	},
	"seg_start": Field{
		Type:        FieldTypeSIZ,
		Name:        "Start",
		Description: "Offset within the LV to the start of the segment in current units.",
	},
	"seg_start_pe": Field{
		Type:        FieldTypeNUM,
		Name:        "Start",
		Description: "Offset within the LV to the start of the segment in physical extents.",
	},
	"seg_size": Field{
		Type:        FieldTypeSIZ,
		Name:        "SSize",
		Description: "Size of segment in current units.",
	},
	"seg_size_pe": Field{
		Type:        FieldTypeSIZ,
		Name:        "SSize",
		Description: "Size of segment in physical extents.",
	},
	"seg_tags": Field{
		Type:        FieldTypeSTR_LIST,
		Name:        "Seg Tags",
		Description: "Tags, if any.",
	},
	"seg_pe_ranges": Field{
		Type:        FieldTypeSTR_LIST,
		Name:        "PE Ranges",
		Description: "Ranges of Physical Extents of underlying devices in command line format (deprecated, use seg_le_ranges for common format).",
	},
	"seg_le_ranges": Field{
		Type:        FieldTypeSTR_LIST,
		Name:        "LE Ranges",
		Description: "Ranges of Logical Extents of underlying devices in command line format.",
	},
	"seg_metadata_le_ranges": Field{
		Type:        FieldTypeSTR_LIST,
		Name:        "Metadata LE Ranges",
		Description: "Ranges of Logical Extents of underlying metadata devices in command line format.",
	},
	"devices": Field{
		Type:        FieldTypeSTR_LIST,
		Name:        "Devices",
		Description: "Underlying devices used with starting extent numbers.",
	},
	"metadata_devices": Field{
		Type:        FieldTypeSTR_LIST,
		Name:        "Metadata Devs",
		Description: "Underlying metadata devices used with starting extent numbers.",
	},
	"seg_monitor": Field{
		Type:        FieldTypeSTR,
		Name:        "Monitor",
		Description: "Dmeventd monitoring status of the segment.",
	},
	"cache_policy": Field{
		Type:        FieldTypeSTR,
		Name:        "CachePolicy",
		Description: "The cache policy (cached segments only).",
	},
	"cache_settings": Field{
		Type:        FieldTypeSTR_LIST,
		Name:        "CacheSettings",
		Description: "Cache settings/parameters (cached segments only).",
	},
	"vdo_compression": Field{
		Type:        FieldTypeBIN,
		Name:        "VDOCompression",
		Description: "Set for compressed LV (vdopool).",
	},
	"vdo_deduplication": Field{
		Type:        FieldTypeBIN,
		Name:        "VDODeduplication",
		Description: "Set for deduplicated LV (vdopool).",
	},
	"vdo_use_metadata_hints": Field{
		Type:        FieldTypeBIN,
		Name:        "VDOMetadataHints",
		Description: "Use REQ_SYNC for writes (vdopool).",
	},
	"vdo_minimum_io_size": Field{
		Type:        FieldTypeNUM,
		Name:        "VDOMinimumIOSize",
		Description: "Minimum acceptable IO size (vdopool).",
	},
	"vdo_block_map_cache_size": Field{
		Type:        FieldTypeSIZ,
		Name:        "VDOBlockMapCacheSize",
		Description: "Allocated caching size (vdopool).",
	},
	"vdo_block_map_era_length": Field{
		Type:        FieldTypeNUM,
		Name:        "VDOBlockMapEraLength",
		Description: "Speed of cache writes (vdopool).",
	},
	"vdo_use_sparse_index": Field{
		Type:        FieldTypeBIN,
		Name:        "VDOSparseIndex",
		Description: "Sparse indexing (vdopool).",
	},
	"vdo_index_memory_size": Field{
		Type:        FieldTypeSIZ,
		Name:        "VDOIndexMemorySize",
		Description: "Allocated indexing memory (vdopool).",
	},
	"vdo_slab_size": Field{
		Type:        FieldTypeSIZ,
		Name:        "VDOSlabSize",
		Description: "Increment size for growing (vdopool).",
	},
	"vdo_ack_threads": Field{
		Type:        FieldTypeNUM,
		Name:        "VDOAckThreads",
		Description: "Acknowledging threads (vdopool).",
	},
	"vdo_bio_threads": Field{
		Type:        FieldTypeNUM,
		Name:        "VDOBioThreads",
		Description: "IO submitting threads (vdopool).",
	},
	"vdo_bio_rotation": Field{
		Type:        FieldTypeNUM,
		Name:        "VDOBioRotation",
		Description: "IO enqueue (vdopool).",
	},
	"vdo_cpu_threads": Field{
		Type:        FieldTypeNUM,
		Name:        "VDOCPUThreads",
		Description: "CPU threads for compression and hashing (vdopool).",
	},
	"vdo_hash_zone_threads": Field{
		Type:        FieldTypeNUM,
		Name:        "VDOHashZoneThreads",
		Description: "Threads for subdivide parts (vdopool).",
	},
	"vdo_logical_threads": Field{
		Type:        FieldTypeNUM,
		Name:        "VDOLogicalThreads",
		Description: "Logical threads for subdivide parts (vdopool).",
	},
	"vdo_physical_threads": Field{
		Type:        FieldTypeNUM,
		Name:        "VDOPhysicalThreads",
		Description: "Physical threads for subdivide parts (vdopool).",
	},
	"vdo_max_discard": Field{
		Type:        FieldTypeNUM,
		Name:        "VDOMaxDiscard",
		Description: "Maximum discard size volume can recieve (vdopool).",
	},
	"vdo_write_policy": Field{
		Type:        FieldTypeSTR,
		Name:        "VDOWritePolicy",
		Description: "Specified write policy (vdopool).",
	},
	"vdo_header_size": Field{
		Type:        FieldTypeSIZ,
		Name:        "VDOHeaderSize",
		Description: "Header size at front of vdopool.",
	},
	"pvseg_start": Field{
		Type:        FieldTypeNUM,
		Name:        "Start",
		Description: "Physical Extent number of start of segment.",
	},
	"pvseg_size": Field{
		Type:        FieldTypeNUM,
		Name:        "SSize",
		Description: "Number of extents in segment.",
	},
}

var lvmEnums = map[string][]string{
	"lv_permissions": []string{"writeable", "read-only", "read-only-override"},
	"lv_when_full":   []string{"error", "queue"},
	"vg_permissions": []string{"writeable", "read-only"},
}

