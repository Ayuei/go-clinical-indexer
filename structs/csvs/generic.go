package csvs

import "time"

type BioRedditSubmissions struct {
	Subreddit string `csv:"subreddit,omitempty"`
	Selftext string `csv:"selftext,omitempty"`
	Author_fullname string `csv:"author_fullname,omitempty"`
	Title string `csv:"title,omitempty"`
	Pwls int `csv:"pwls,omitempty"`
	Name string `csv:"name,omitempty"`
	Upvote_ratio float32 `csv:"upvote_ratio,omitempty"`
	Author_flair_background_color string `csv:"author_flair_background_color,omitempty"`
	Ups int `csv:"ups,omitempty"`
	Total_awards_received int `csv:"total_awards_received,omitempty"`
	Author_flair_template_id string `csv:"author_flair_template_id,omitempty"`
	User_reports string `csv:"user_reports,omitempty"`
	Link_flair_text string `csv:"link_flair_text,omitempty"`
	Score int `csv:"score,omitempty"`
	Approved_by string `csv:"approved_by,omitempty"`
	Thumbnail string `csv:"thumbnail,omitempty"`
	Edited time.Time `csv:"edited,omitempty"`
	Author_flair_css_class string `csv:"author_flair_css_class,omitempty"`
	Gildings string `csv:"gildings,omitempty"`
	Content_categories string `csv:"content_categories,omitempty"`
	Is_self bool `csv:"is_self,omitempty"`
	Mod_note string `csv:"mod_note,omitempty"`
	Created time.Time `csv:"created,omitempty"`
	Wls int`csv:"wls,omitempty"`
	Removed_by_category string `csv:"removed_by_category,omitempty"`
	Banned_by string `csv:"banned_by,omitempty"`
	Allow_live_comments bool`csv:"allow_live_comments,omitempty"`
	Selftext_html string `csv:"selftext_html,omitempty"`
	Likes int`csv:"likes,omitempty"`
	Banned_at_utc time.Time `csv:"banned_at_utc,omitempty"`
	No_follow string `csv:"no_follow,omitempty"`
	Is_crosspostable string `csv:"is_crosspostable,omitempty"`
	Pinned bool `csv:"pinned,omitempty"`
	Over_18 bool `csv:"over_18,omitempty"`
	All_awardings string `csv:"all_awardings,omitempty"`
	Awarders string `csv:"awarders,omitempty"`
	Can_gild bool `csv:"can_gild,omitempty"`
	Spoiler bool `csv:"spoiler,omitempty"`
	Locked bool `csv:"locked,omitempty"`
	Author_flair_text string `csv:"author_flair_text,omitempty"`
	Treatment_tags string `csv:"treatment_tags,omitempty"`
	Visited string `csv:"visited,omitempty"`
	Removed_by string `csv:"removed_by,omitempty"`
	Num_reports string `csv:"num_reports,omitempty"`
	Distinguished string `csv:"distinguished,omitempty"`
	Subreddit_id string `csv:"subreddit_id,omitempty"`
	Author_is_blocked bool `csv:"author_is_blocked,omitempty"`
	Mod_reason_by string `csv:"mod_reason_by,omitempty"`
	Removal_reason string `csv:"removal_reason,omitempty"`
	Id string `csv:"id,omitempty"`
	Is_robot_indexable bool `csv:"is_robot_indexable,omitempty"`
	Report_reasons string `csv:"report_reasons,omitempty"`
	Author string `csv:"author,omitempty"`
	Discussion_type string `csv:"discussion_type,omitempty"`
	Num_comments int `csv:"num_comments,omitempty"`
	Whitelist_status string `csv:"whitelist_status,omitempty"`
	Mod_reports string `csv:"mod_reports,omitempty"`
	Permalink string `csv:"permalink,omitempty"`
	Stickied bool `csv:"stickied,omitempty"`
	Url string `csv:"url,omitempty"`
	Created_utc time.Time`csv:"created_utc,omitempty"`
	Num_crossposts int `csv:"num_crossposts,omitempty"`
	Media string `csv:"media,omitempty"`
	Is_video bool `csv:"is_video,omitempty"`
	Post_hint string `csv:"post_hint,omitempty"`
	Preview string `csv:"preview,omitempty"`
	Link_flair_template_id string `csv:"link_flair_template_id,omitempty"`
	Crosspost_parent_list string `csv:"crosspost_parent_list,omitempty"`
	Url_overridden_by_dest bool `csv:"url_overridden_by_dest,omitempty"`
	Crosspost_parent string `csv:"crosspost_parent,omitempty"`
	Media_metadata string `csv:"media_metadata,omitempty"`
}

type BioRedditComments struct {
	Subreddit_id string `csv:"subreddit_id,omitempty"`
	Author_is_blocked bool `csv:"author_is_blocked,omitempty"`
	Edited time.Time `csv:"edited,omitempty"`
	Mod_reason_by string `csv:"mod_reason_by,omitempty"`
	Banned_by string `csv:"banned_by,omitempty"`
	Ups int `csv:"ups,omitempty"`
	Num_reports int `csv:"num_reports,omitempty"`
	Total_awards_received int `csv:"total_awards_received,omitempty"`
	Subreddit string `csv:"subreddit,omitempty"`
	Author_flair_template_id string `csv:"author_flair_template_id,omitempty"`
	Likes int `csv:"likes,omitempty"`
	User_reports string `csv:"user_reports,omitempty"`
	Id string `csv:"id,omitempty"`
	Banned_at_utc time.Time `csv:"banned_at_utc,omitempty"`
	Mod_reason_title string `csv:"mod_reason_title,omitempty"`
	Gilded int `csv:"gilded,omitempty"`
	Author string `csv:"author,omitempty"`
	Parent_id string `csv:"parent_id,omitempty"`
	Score int `csv:"score,omitempty"`
	Author_fullname string `csv:"author_fullname,omitempty"`
	Report_reasons string `csv:"report_reasons,omitempty"`
	Removal_reason string `csv:"removal_reason,omitempty"`
	Approved_by string `csv:"approved_by,omitempty"`
	All_awardings string `csv:"all_awardings,omitempty"`
	Body string `csv:"body,omitempty"`
	Awarders string `csv:"awarders,omitempty"`
	Top_awarded_type string `csv:"top_awarded_type,omitempty"`
	Downs int `csv:"downs,omitempty"`
	Author_flair_css_class string `csv:"author_flair_css_class,omitempty"`
	Author_patreon_flair bool `csv:"author_patreon_flair,omitempty"`
	Author_flair_richtext string `csv:"author_flair_richtext,omitempty"`
	Body_html string `csv:"body_html,omitempty"`
	Gildings string `csv:"gildings,omitempty"`
	Stickied bool `csv:"stickied,omitempty"`
	Author_premium bool `csv:"author_premium,omitempty"`
	Link_id string `csv:"link_id,omitempty"`
	Score_hidden bool `csv:"score_hidden,omitempty"`
	Permalink string `csv:"permalink,omitempty"`
	Locked bool `csv:"locked,omitempty"`
	Name string `csv:"name,omitempty"`
	Created time.Time `csv:"created,omitempty"`
	Author_flair_text string `csv:"author_flair_text,omitempty"`
	Treatment_tags string `csv:"treatment_tags,omitempty"`
	Created_utc time.Time `csv:"created_utc,omitempty"`
	Controversiality int `csv:"controversiality,omitempty"`
	Mod_reports string `csv:"mod_reports,omitempty"`
	Mod_note string `csv:"mod_note,omitempty"`
	Distinguished string `csv:"distinguished,omitempty"`
}
