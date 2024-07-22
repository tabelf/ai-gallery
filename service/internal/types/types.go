// Code generated by goctl. DO NOT EDIT.
package types

type AddUserRequest struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Role     string `json:"role,options=ADMIN|USER"`
}

type AnalysisTaskBo struct {
	Name string `json:"name"`
	Data []int  `json:"data"`
}

type AnalysisUserBo struct {
	AuthorID   int    `json:"author_id"`
	AuthorName string `json:"author_name"`
	Count      int    `json:"count"`
}

type AuthLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthLoginResponse struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

type CreateSubTaskReq struct {
	TaskID   int64  `json:"task_id"`
	Base64   string `json:"base64"`
	Filename string `json:"filename"`
}

type CreateSubTaskResp struct {
	SubTaskID int `json:"sub_task_id"`
}

type CreateTaskReq struct {
	Category       string      `json:"category"`
	Prompt         string      `json:"prompt"`
	NegativePrompt string      `json:"negative_prompt"`
	RefImages      []TaskImage `json:"ref_images"`
	Width          float32     `json:"width"`
	Height         float32     `json:"height"`
	Seed           string      `json:"seed"`
	SamplerName    string      `json:"sampler_name"`
	Steps          int         `json:"steps"` // 采样步数
	CfgScale       int         `json:"cfg_scale"`
	BatchSize      int         `json:"batch_size"`
	Total          int         `json:"total"`
	SdModelName    string      `json:"sd_model_name"`
	SdModelHash    string      `json:"sd_model_hash"`
	SdVaeName      *string     `json:"sd_vae_name,optional"`
	SdVaeHash      *string     `json:"sd_vae_hash,optional"`
	JobTimestamp   string      `json:"job_timestamp"`
	Version        string      `json:"version"`
}

type CreateTaskResp struct {
	TaskID int `json:"task_id"`
}

type DeleteUserRequest struct {
	UserID int `path:"user_id"`
}

type GetAnalysisBaseResponse struct {
	TaskTotal           int `json:"task_total"`            // 任务总数
	WorkTotal           int `json:"work_total"`            // 作品总数
	ExcellentTotal      int `json:"excellent_total"`       // 优秀作品数
	AccountTotal        int `json:"account_total"`         // 用户总数
	TodayTaskCount      int `json:"today_task_count"`      // 今日任务数
	TodayWorkCount      int `json:"today_work_count"`      // 今日作品总数
	TodayExcellentCount int `json:"today_excellent_count"` // 今日优秀作品数
}

type GetAnalysisTaskRequest struct {
	Day int `form:"day,default=7"` // 几天内
}

type GetAnalysisTaskResponse struct {
	Times          []string          `json:"times"`
	AnalysisTaskBo []*AnalysisTaskBo `json:"tasks"`
}

type GetAnalysisUserRequest struct {
	Day int `form:"day,default=1"` // 几天内
}

type GetAnalysisUserResponse struct {
	AnalysisUserBo []*AnalysisUserBo `json:"users"`
}

type GetSettingResponse struct {
	StoreName    string `json:"store_name"`
	StoreAddress string `json:"store_address"`
	SecureID     string `json:"secure_id"`
	SecureKey    string `json:"secure_key"`
	StoreBucket  string `json:"store_bucket"`
	InitPwd      string `json:"init_pwd"`
}

type GetTaskAuthorResponse struct {
	TaskAuthorBo []*TaskAuthorBo `json:"data"`
}

type GetTaskModelResponse struct {
	Models []string `json:"models"`
}

type GetTaskRequest struct {
	Offset    int    `form:"offset,default=0"`
	Limit     int    `form:"limit,default=10"`
	StartTime string `form:"start_time,optional"`
	EndTime   string `form:"end_time,optional"`
	AuthorIDs []int  `form:"author_ids,optional"`
	Category  string `form:"category,optional"`
	Status    string `form:"status,optional"`
}

type GetTaskResponse struct {
	Total int       `json:"total"`
	Data  []*TaskBo `json:"data"`
}

type GetTaskSizeResponse struct {
	Sizes []string `json:"sizes"`
}

type GetUserDetailResponse struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Role     string `json:"role"`
}

type GetUserRequest struct {
	Offset   int    `form:"offset,default=0"`
	Limit    int    `form:"limit,default=10"`
	Username string `form:"username,optional"`
	Nickname string `form:"nickname,optional"`
	Role     string `form:"role,options=ADMIN|USER,optional"`
	Status   int    `form:"status,default=-1,optional"`
}

type GetUserResponse struct {
	Total int       `json:"total"`
	Data  []*UserBo `json:"data"`
}

type GetWorkDetailRequest struct {
	TaskID int `path:"task_id"`
}

type GetWorkDetailResponse struct {
	TaskID         int             `json:"task_id"`
	Category       string          `json:"category"`
	Prompt         string          `json:"prompt"`
	NegativePrompt string          `json:"negative_prompt"`
	Width          float32         `json:"width"`
	Height         float32         `json:"height"`
	Size           string          `json:"size"`
	Seed           string          `json:"seed"`
	SamplerName    string          `json:"sampler_name"`
	Steps          int             `json:"steps"` // 采样步数
	CfgScale       int             `json:"cfg_scale"`
	BatchSize      int             `json:"batch_size"`
	Total          int             `json:"total"`
	SdModelName    string          `json:"sd_model_name"`
	SdModelHash    string          `json:"sd_model_hash"`
	SdVaeName      string          `json:"sd_vae_name"`
	SdVaeHash      string          `json:"sd_vae_hash"`
	JobTimestamp   string          `json:"job_timestamp"`
	Version        string          `json:"version"`
	AuthorID       int             `json:"author_id"`
	AuthorName     string          `json:"author_name"`
	RefImages      []string        `json:"ref_images"`
	WorkDetailBos  []*WorkDetailBo `json:"details"`
}

type GetWorkRequest struct {
	Offset       int    `form:"offset,default=0"`
	Limit        int    `form:"limit,default=10"`
	Sorted       string `form:"sorted,default=asc,optional"`
	Size         string `form:"size,optional"`
	SdModelName  string `json:"sd_model_name,optional"`
	HasExcellent *int   `form:"has_excellent,default=-1,optional"`
	HasRefImage  *int   `form:"has_ref_image,default=-1,optional"`
}

type GetWorkResponse struct {
	Total int       `json:"total"`
	Data  []*WorkBo `json:"data"`
}

type MarkWorkExcellentRequest struct {
	TaskID       int  `json:"task_id"`
	SubTaskID    int  `json:"sub_task_id"`
	HasExcellent bool `json:"has_excellent"`
}

type PingResp struct {
	Message string `json:"message"`
}

type ResetPwdRequest struct {
	UserID int `path:"user_id"`
}

type TaskAuthorBo struct {
	AuthorID   int    `json:"author_id"`
	AuthorName string `json:"author_name"`
}

type TaskBo struct {
	TaskID       int      `json:"task_id"`
	Category     string   `json:"category"`
	Prompt       string   `json:"prompt"`
	Size         string   `json:"size"`
	AuthorID     int      `json:"author_id"`
	AuthorName   string   `json:"author_name"`
	Total        int      `json:"total"`
	JobTimestamp string   `json:"job_timestamp"`
	Status       string   `json:"status"`
	Count        int      `json:"count"`
	ImageUrls    []string `json:"image_urls"`
}

type TaskImage struct {
	Base64   string `json:"base64"`
	Filename string `json:"filename"`
}

type UpdateSettingRequest struct {
	StoreName    string `json:"store_name"`
	StoreAddress string `json:"store_address"`
	SecureID     string `json:"secure_id,optional"`
	SecureKey    string `json:"secure_key,optional"`
	StoreBucket  string `json:"store_bucket,optional"`
	InitPwd      string `json:"init_pwd"`
}

type UpdateTaskReq struct {
	TaskID   int    `path:"task_id"`
	Base64   string `json:"base64,optional"`
	Filename string `json:"filename,optional"`
	Extra    string `json:"extra,optional"`
}

type UpdateUserDetailRequest struct {
	UserID      int    `json:"user_id"`
	Username    string `json:"username"`
	Nickname    string `json:"nickname"`
	OldPassword string `json:"old_password,optional"`
	NewPassword string `json:"new_password,optional"`
}

type UpdateUserRequest struct {
	UserID   int    `path:"user_id"`
	Username string `json:"username,optional"`
	Nickname string `json:"nickname,optional"`
	Role     string `json:"role,options=ADMIN|USER,optional"`
	Status   int    `json:"status,default=-1,optional"`
}

type UserBo struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Role     string `json:"role"`
	Status   int    `json:"status"`
}

type WorkBo struct {
	TaskID     int    `json:"task_id"`
	Category   string `json:"category"`
	Prompt     string `json:"prompt"`
	Size       string `json:"size"`
	AuthorID   int    `json:"author_id"`
	AuthorName string `json:"author_name"`
	Count      int    `json:"count"`
	HeadImage  string `json:"head_image"`
}

type WorkDetailBo struct {
	SubTaskID    int    `json:"sub_task_id"`
	ImageURL     string `json:"image_url"`
	HasExcellent bool   `json:"has_excellent"`
}