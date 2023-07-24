package postgres

const (
	tableMasterChoreo       = "m_choreo"
	tableMasterChoreoDetail = "m_det_choreo"

	columnSelectAllChoreo       = "choreo_id,title,description,difficulty,duration,is_active,position,vid_preview_url,vid_thumbnail_url,vid_preview_url_cdn,vid_thumbnail_url_cdn,choreographer_id,music_id,additional_info,temp_price"
	columnSelectAllChoreoDetail = "det_choreo_id,choreo_id,title,duration,is_active,position,vid_url,vid_thumbnail_url,vid_url_cdn,vid_thumbnail_url_cdn,vid_test_url,vid_test_url_cdn,vision_body_pose,vision_angle_threshold,vision_time_offset"

	columnInsertChoreo       = "title,description,difficulty,duration,is_active,position,vid_preview_url_cdn,vid_thumbnail_url_cdn,choreographer_id,music_id,additional_info,temp_price"
	columnInsertChoreoDetail = "choreo_id,title,duration,is_active,position,vid_url_cdn,vid_thumbnail_url_cdn,vid_test_url_cdn,vision_body_pose,vision_angle_threshold,vision_time_offset"
)
