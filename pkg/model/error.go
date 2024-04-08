package models


type Error struct {
    ErrorCode   int    `json:"error_code"`
    ErrorMessage string `json:"error_message"`
    Details      string `json:"details,omitempty"`
}