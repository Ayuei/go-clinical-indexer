package csv

import "time"

type BioRedditSubmissions struct {
	subreddit string `csv:"subreddit,omitempty"`
	selftext string `csv:"selftext,omitempty"`
	author_fullname string `csv:"author_fullname,omitempty"`
	title string `csv:"title,omitempty"`
	pwls int `csv:"pwls,omitempty"`
	name string `csv:"name,omitempty"`
	upvote_ratio float32 `csv:"upvote_ratio,omitempty"`
	author_flair_background_color string `csv:"author_flair_background_color,omitempty"`
	ups int `csv:"ups,omitempty"`
	total_awards_received int `csv:"total_awards_received,omitempty"`
	author_flair_template_id string `csv:"author_flair_template_id,omitempty"`
	user_reports string `csv:"user_reports,omitempty"`
	link_flair_text string `csv:"link_flair_text,omitempty"`
	score int `csv:"score,omitempty"`
	approved_by string `csv:"approved_by,omitempty"`
	thumbnail string `csv:"thumbnail,omitempty"`
	edited time.Time `csv:"edited,omitempty"`
	author_flair_css_class string `csv:"author_flair_css_class,omitempty"`
	gildings string `csv:"gildings,omitempty"`
	content_categories string `csv:"content_categories,omitempty"`
	is_self bool `csv:"is_self,omitempty"`
	mod_note string `csv:"mod_note,omitempty"`
	created time.Time `csv:"created,omitempty"`
	wls int`csv:"wls,omitempty"`
	removed_by_category string `csv:"removed_by_category,omitempty"`
	banned_by string `csv:"banned_by,omitempty"`
	allow_live_comments bool`csv:"allow_live_comments,omitempty"`
	selftext_html string `csv:"selftext_html,omitempty"`
	likes int`csv:"likes,omitempty"`
	banned_at_utc time.Time `csv:"banned_at_utc,omitempty"`
	no_follow string `csv:"no_follow,omitempty"`
	is_crosspostable string `csv:"is_crosspostable,omitempty"`
	pinned bool `csv:"pinned,omitempty"`
	over_18 bool `csv:"over_18,omitempty"`
	all_awardings string `csv:"all_awardings,omitempty"`
	awarders string `csv:"awarders,omitempty"`
	can_gild bool `csv:"can_gild,omitempty"`
	spoiler bool `csv:"spoiler,omitempty"`
	locked bool `csv:"locked,omitempty"`
	author_flair_text string `csv:"author_flair_text,omitempty"`
	treatment_tags string `csv:"treatment_tags,omitempty"`
	visited string `csv:"visited,omitempty"`
	removed_by string `csv:"removed_by,omitempty"`
	num_reports string `csv:"num_reports,omitempty"`
	distinguished string `csv:"distinguished,omitempty"`
	subreddit_id string `csv:"subreddit_id,omitempty"`
	author_is_blocked bool `csv:"author_is_blocked,omitempty"`
	mod_reason_by string `csv:"mod_reason_by,omitempty"`
	removal_reason string `csv:"removal_reason,omitempty"`
	id string `csv:"id,omitempty"`
	is_robot_indexable bool `csv:"is_robot_indexable,omitempty"`
	report_reasons string `csv:"report_reasons,omitempty"`
	author string `csv:"author,omitempty"`
	discussion_type string `csv:"discussion_type,omitempty"`
	num_comments int `csv:"num_comments,omitempty"`
	whitelist_status string `csv:"whitelist_status,omitempty"`
	mod_reports string `csv:"mod_reports,omitempty"`
	permalink string `csv:"permalink,omitempty"`
	stickied bool `csv:"stickied,omitempty"`
	url string `csv:"url,omitempty"`
	created_utc time.Time`csv:"created_utc,omitempty"`
	num_crossposts int `csv:"num_crossposts,omitempty"`
	media string `csv:"media,omitempty"`
	is_video bool `csv:"is_video,omitempty"`
	post_hint string `csv:"post_hint,omitempty"`
	preview string `csv:"preview,omitempty"`
	link_flair_template_id string `csv:"link_flair_template_id,omitempty"`
	crosspost_parent_list string `csv:"crosspost_parent_list,omitempty"`
	url_overridden_by_dest bool `csv:"url_overridden_by_dest,omitempty"`
	crosspost_parent string `csv:"crosspost_parent,omitempty"`
	media_metadata string `csv:"media_metadata,omitempty"`
}

type BioRedditComments struct {
	subreddit_id string `csv:"subreddit_id,omitempty"`
	author_is_blocked bool `csv:"author_is_blocked,omitempty"`
	edited time.Time `csv:"edited,omitempty"`
	mod_reason_by string `csv:"mod_reason_by,omitempty"`
	banned_by string `csv:"banned_by,omitempty"`
	ups int `csv:"ups,omitempty"`
	num_reports int `csv:"num_reports,omitempty"`
	total_awards_received int `csv:"total_awards_received,omitempty"`
	subreddit string `csv:"subreddit,omitempty"`
	author_flair_template_id string `csv:"author_flair_template_id,omitempty"`
	likes int `csv:"likes,omitempty"`
	user_reports string `csv:"user_reports,omitempty"`
	id string `csv:"id,omitempty"`
	banned_at_utc time.Time `csv:"banned_at_utc,omitempty"`
	mod_reason_title string `csv:"mod_reason_title,omitempty"`
	gilded int `csv:"gilded,omitempty"`
	author string `csv:"author,omitempty"`
	parent_id string `csv:"parent_id,omitempty"`
	score int `csv:"score,omitempty"`
	author_fullname string `csv:"author_fullname,omitempty"`
	report_reasons string `csv:"report_reasons,omitempty"`
	removal_reason string `csv:"removal_reason,omitempty"`
	approved_by string `csv:"approved_by,omitempty"`
	all_awardings string `csv:"all_awardings,omitempty"`
	body string `csv:"body,omitempty"`
	awarders string `csv:"awarders,omitempty"`
	top_awarded_type string `csv:"top_awarded_type,omitempty"`
	downs int `csv:"downs,omitempty"`
	author_flair_css_class string `csv:"author_flair_css_class,omitempty"`
	author_patreon_flair bool `csv:"author_patreon_flair,omitempty"`
	author_flair_richtext string `csv:"author_flair_richtext,omitempty"`
	body_html string `csv:"body_html,omitempty"`
	gildings string `csv:"gildings,omitempty"`
	stickied bool `csv:"stickied,omitempty"`
	author_premium bool `csv:"author_premium,omitempty"`
	link_id string `csv:"link_id,omitempty"`
	score_hidden bool `csv:"score_hidden,omitempty"`
	permalink string `csv:"permalink,omitempty"`
	locked bool `csv:"locked,omitempty"`
	name string `csv:"name,omitempty"`
	created time.Time `csv:"created,omitempty"`
	author_flair_text string `csv:"author_flair_text,omitempty"`
	treatment_tags string `csv:"treatment_tags,omitempty"`
	created_utc time.Time `csv:"created_utc,omitempty"`
	controversiality int `csv:"controversiality,omitempty"`
	mod_reports string `csv:"mod_reports,omitempty"`
	mod_note string `csv:"mod_note,omitempty"`
	distinguished string `csv:"distinguished,omitempty"`
}
