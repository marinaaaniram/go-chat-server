package converter

// import (
// 	desc "github.com/marinaaaniram/go-chat-server/pkg/chat_v1"
// )

// Convert Chat internal model to desc model
// func FromChatIdToDescCreate(id int64) *desc.CreateResponse {
// 	return &desc.CreateResponse{
// 		Id: id,
// 	}
// }

// Convert desc CreateRequest fields to internal Chat model
// func FromDescCreateToChat(req *desc.CreateRequest) *model.Chat {
// 	if req == nil {
// 		return nil
// 	}

// 	return &model.Chat{
// 		Usernames: req.GetUsernames(),
// 	}
// }

// Convert desc DeleteRequest fields to internal Chat model
// func FromDescDeleteToChat(req *desc.DeleteRequest) *model.Chat {
// 	if req == nil {
// 		return nil
// 	}

// 	return &model.Chat{
// 		ID: req.GetId(),
// 	}
// }
