package helper

import "jamal/api/models/web"

func HandleErrorResponse(response *web.ProductResponse, err error) {
	if err != nil {
		// RESPONE KTIKA INTERNAL SERVER RROR
		if response.Code != 404 {
			*response = web.ProductResponse{
				Code:   500,
				Status: "Internal Server Error",
				Data:   nil,
			}
		}
	}
}
