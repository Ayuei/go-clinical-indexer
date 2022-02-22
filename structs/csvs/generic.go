package csvs

import "time"

type BioRedditSubmissions struct {
	subreddit string `csvs:"subreddit,omitempty"`
	selftext string `csvs:"selftext,omitempty"`
	author_fullname string `csvs:"author_fullname,omitempty"`
	title string `csvs:"title,omitempty"`
	pwls int `csvs:"pwls,omitempty"`
	name string `csvs:"name,omitempty"`
	upvote_ratio float32 `csvs:"upvote_ratio,omitempty"`
	author_flair_background_color string `csvs:"author_flair_background_color,omitempty"`
	ups int `csvs:"ups,omitempty"`
	total_awards_received int `csvs:"total_awards_received,omitempty"`
	author_flair_template_id string `csvs:"author_flair_template_id,omitempty"`
	user_reports string `csvs:"user_reports,omitempty"`
	link_flair_text string `csvs:"link_flair_text,omitempty"`
	score int `csvs:"score,omitempty"`
	approved_by string `csvs:"approved_by,omitempty"`
	thumbnail string `csvs:"thumbnail,omitempty"`
	edited time.Time `csvs:"edited,omitempty"`
	author_flair_css_class string `csvs:"author_flair_css_class,omitempty"`
	gildings string `csvs:"gildings,omitempty"`
	content_categories string `csvs:"content_categories,omitempty"`
	is_self bool `csvs:"is_self,omitempty"`
	mod_note string `csvs:"mod_note,omitempty"`
	created time.Time `csvs:"created,omitempty"`
	wls int`csvs:"wls,omitempty"`
	removed_by_category string `csvs:"removed_by_category,omitempty"`
	banned_by string `csvs:"banned_by,omitempty"`
	allow_live_comments bool`csvs:"allow_live_comments,omitempty"`
	selftext_html string `csvs:"selftext_html,omitempty"`
	likes int`csvs:"likes,omitempty"`
	banned_at_utc time.Time `csvs:"banned_at_utc,omitempty"`
	no_follow string `csvs:"no_follow,omitempty"`
	is_crosspostable string `csvs:"is_crosspostable,omitempty"`
	pinned bool `csvs:"pinned,omitempty"`
	over_18 bool `csvs:"over_18,omitempty"`
	all_awardings string `csvs:"all_awardings,omitempty"`
	awarders string `csvs:"awarders,omitempty"`
	can_gild bool `csvs:"can_gild,omitempty"`
	spoiler bool `csvs:"spoiler,omitempty"`
	locked bool `csvs:"locked,omitempty"`
	author_flair_text string `csvs:"author_flair_text,omitempty"`
	treatment_tags string `csvs:"treatment_tags,omitempty"`
	visited string `csvs:"visited,omitempty"`
	removed_by string `csvs:"removed_by,omitempty"`
	num_reports string `csvs:"num_reports,omitempty"`
	distinguished string `csvs:"distinguished,omitempty"`
	subreddit_id string `csvs:"subreddit_id,omitempty"`
	author_is_blocked bool `csvs:"author_is_blocked,omitempty"`
	mod_reason_by string `csvs:"mod_reason_by,omitempty"`
	removal_reason string `csvs:"removal_reason,omitempty"`
	id string `csvs:"id,omitempty"`
	is_robot_indexable bool `csvs:"is_robot_indexable,omitempty"`
	report_reasons string `csvs:"report_reasons,omitempty"`
	author string `csvs:"author,omitempty"`
	discussion_type string `csvs:"discussion_type,omitempty"`
	num_comments int `csvs:"num_comments,omitempty"`
	whitelist_status string `csvs:"whitelist_status,omitempty"`
	mod_reports string `csvs:"mod_reports,omitempty"`
	permalink string `csvs:"permalink,omitempty"`
	stickied bool `csvs:"stickied,omitempty"`
	url string `csvs:"url,omitempty"`
	created_utc time.Time`csvs:"created_utc,omitempty"`
	num_crossposts int `csvs:"num_crossposts,omitempty"`
	media string `csvs:"media,omitempty"`
	is_video bool `csvs:"is_video,omitempty"`
	post_hint string `csvs:"post_hint,omitempty"`
	preview string `csvs:"preview,omitempty"`
	link_flair_template_id string `csvs:"link_flair_template_id,omitempty"`
	crosspost_parent_list string `csvs:"crosspost_parent_list,omitempty"`
	url_overridden_by_dest bool `csvs:"url_overridden_by_dest,omitempty"`
	crosspost_parent string `csvs:"crosspost_parent,omitempty"`
	media_metadata string `csvs:"media_metadata,omitempty"`
}

type BioRedditComments struct {
	subreddit_id string `csvs:"subreddit_id,omitempty"`
	author_is_blocked bool `csvs:"author_is_blocked,omitempty"`
	edited time.Time `csvs:"edited,omitempty"`
	mod_reason_by string `csvs:"mod_reason_by,omitempty"`
	banned_by string `csvs:"banned_by,omitempty"`
	ups int `csvs:"ups,omitempty"`
	num_reports int `csvs:"num_reports,omitempty"`
	total_awards_received int `csvs:"total_awards_received,omitempty"`
	subreddit string `csvs:"subreddit,omitempty"`
	author_flair_template_id string `csvs:"author_flair_template_id,omitempty"`
	likes int `csvs:"likes,omitempty"`
	user_reports string `csvs:"user_reports,omitempty"`
	id string `csvs:"id,omitempty"`
	banned_at_utc time.Time `csvs:"banned_at_utc,omitempty"`
	mod_reason_title string `csvs:"mod_reason_title,omitempty"`
	gilded int `csvs:"gilded,omitempty"`
	author string `csvs:"author,omitempty"`
	parent_id string `csvs:"parent_id,omitempty"`
	score int `csvs:"score,omitempty"`
	author_fullname string `csvs:"author_fullname,omitempty"`
	report_reasons string `csvs:"report_reasons,omitempty"`
	removal_reason string `csvs:"removal_reason,omitempty"`
	approved_by string `csvs:"approved_by,omitempty"`
	all_awardings string `csvs:"all_awardings,omitempty"`
	body string `csvs:"body,omitempty"`
	awarders string `csvs:"awarders,omitempty"`
	top_awarded_type string `csvs:"top_awarded_type,omitempty"`
	downs int `csvs:"downs,omitempty"`
	author_flair_css_class string `csvs:"author_flair_css_class,omitempty"`
	author_patreon_flair bool `csvs:"author_patreon_flair,omitempty"`
	author_flair_richtext string `csvs:"author_flair_richtext,omitempty"`
	body_html string `csvs:"body_html,omitempty"`
	gildings string `csvs:"gildings,omitempty"`
	stickied bool `csvs:"stickied,omitempty"`
	author_premium bool `csvs:"author_premium,omitempty"`
	link_id string `csvs:"link_id,omitempty"`
	score_hidden bool `csvs:"score_hidden,omitempty"`
	permalink string `csvs:"permalink,omitempty"`
	locked bool `csvs:"locked,omitempty"`
	name string `csvs:"name,omitempty"`
	created time.Time `csvs:"created,omitempty"`
	author_flair_text string `csvs:"author_flair_text,omitempty"`
	treatment_tags string `csvs:"treatment_tags,omitempty"`
	created_utc time.Time `csvs:"created_utc,omitempty"`
	controversiality int `csvs:"controversiality,omitempty"`
	mod_reports string `csvs:"mod_reports,omitempty"`
	mod_note string `csvs:"mod_note,omitempty"`
	distinguished string `csvs:"distinguished,omitempty"`
}
