package postgres

const (
	tableMasterChoreo       = "m_choreo"
	tableMasterChoreoDetail = "m_det_choreo"

	columnSelectAllChoreo       = "choreo_id,title,description,difficulty,duration,is_active,position,vid_preview_url,vid_thumbnail_url,choreographer_id,music_id,temp_price"
	columnSelectAllChoreoDetail = "det_choreo_id,choreo_id,title,duration,is_active,position,vid_url,vid_thumbnail_url,vision_body_pose,vision_angle_threshold,vision_time_offset"
)
