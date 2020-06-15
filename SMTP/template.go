package SMTP

import (
	"fmt"
	"server/DataModels"
	"time"
)

// func ActionItemSMTPTemplate(actionItem DataModels.ActionItems, momName string) string {
// 	deadLine := time.Unix(actionItem.ExpectedCompletionDate, 0)
// 	year, month, date := deadLine.Date()

// 	return fmt.Sprintf(`
// 		<html>
// 		   <head>
// 			  <style>
// 			  </style>
// 		   </head>
// 		   <body style="margin: 50px; ">
// 			  <div style="box-shadow: 6px 6px 6px 3px rgba(100, 100, 100, 0.1); width: 450px; font-family: Lucida Sans Unicode, Lucida Grande, sans-serif;">
// 				 <div style="background-color: #2196F3; width: 410px; color: white; padding: 20px;">
// 					<h4 style="font-weight: lighter;"> <b>Task :</b> %s </h4>
// 				 </div>
// 				 <div
// 					style=" background-color: white; width: 400px; color: black; padding: 20px; font-size: smaller;">
// 					<p>Meeting Name: %s </p>
// 					<p>Target Date: %s </p>
// 					<p>Priority: %s </p>
// 					<p>Assignee: %s </p>
// 					<p>Assigner: %s </p>
// 				 </div>
// 			  </div>
// 		   </body>
// 		</html>`,

// 		actionItem.Description,
// 		momName,
// 		fmt.Sprintf("%d-%d-%d", date, month, year),
// 		actionItem.Priority,
// 		strings.Join(DataAccess.GetUserNameByID(actionItem.AssignedTo), " ,"),
// 		actionItem.AssignedBy,
// 	)
// }

// func CommentSMTPTemplate(actionItem DataModels.ActionItems, comments []DataModels.Comment, commentBy string) string {
// 	deadLine := time.Unix(actionItem.ExpectedCompletionDate, 0)
// 	year, month, date := deadLine.Date()
// 	commentTemplate := ""

// 	for _, commentInstance := range comments {
// 		commentTime := time.Unix(commentInstance.CommentON, 0)
// 		commentYear, commentMonth, commentDate := commentTime.Date()
// 		hour := commentTime.Hour()
// 		minute := commentTime.Minute()

// 		fmt.Println("comment time", commentTime, hour, minute)

// 		commentTemplate = commentTemplate + fmt.Sprintf(`
// 				<p  style="font-size:large; padding-top: 20px;"> %s </p>
//                 <p style="font-size: smaller; color: grey; line-height: 0px;"> %s - %d-%d-%d %s </p>
// 			`,
// 			commentInstance.Comment,
// 			commentInstance.CommentBY,
// 			commentYear, commentMonth, commentDate, commentTime.Format("03:04 PM"),
// 		)
// 	}

// 	return fmt.Sprintf(`<html>
// 	   <head>
// 		  <style>
// 		  </style>
// 	   </head>
//        <body style="margin: 50px; font-family: Lucida Sans Unicode, Lucida Grande, sans-serif;">
// 		  <div style="width: 450px;">
// 			 <div style="background-color: #2196F3; width: 410px; color: white; padding: 20px;">
//                 <h2 style="font-weight: lighter;">Task - %s </h2>
//                 <p style="font-weight: lighter; font-size: small;">Meeting Name : %s </p>
//                 <p style="font-weight: lighter; font-size: small;">Assigner: %s </p>
// 				<p style="font-weight: lighter; font-size: small;">Assignee: %s </p>
//                 <p style="font-weight: lighter; font-size: small;">Priority: %s </p>
//                 <p style="font-weight: lighter; font-size: small;">Status: %s </p>
// 				<p style="font-weight: lighter; font-size: small;">Target Date: %d-%d-%d </p>
// 				<p style="font-weight: lighter; font-size: small;">(Shared by) : %s</p>
// 			 </div>
// 			 <div
//                 style=" background-color: white; width: 408px; color: black; padding: 20px; font-size: smaller; border: #2196F3 1px solid;"> <br> Discussion `+commentTemplate+`

// 			 </div>
// 		  </div>
// 	   </body>
// 	</html>`,
// 		actionItem.Description,
// 		DataAccess.GetMOMData([]string{actionItem.MomID})[0].Name,
// 		actionItem.AssignedBy,
// 		strings.Join(actionItem.AssignedTo, ","),
// 		actionItem.Priority,
// 		actionItem.Status,
// 		date, month, year,
// 		commentBy,
// 	)
// }

func TodoSMTPTemplate(todo []DataModels.Todo) string {

	todoTemplete := ""

	for _, todoInstanse := range todo {
		todoTime := time.Unix(todoInstanse.TaskDate, 0)

		todoTemplete = todoTemplete + fmt.Sprintf(`
				<li  style="font-size:large; padding-top: 20px;"> time:%s %s </li>
			`,
			todoTime.Format("03:04 PM"),
			todoInstanse.Task,
		)
	}

	return fmt.Sprintf(`<html>
	<head>
	   <style>
	   </style>
	</head>
	<body style="margin: 50px; font-family: Lucida Sans Unicode, Lucida Grande, sans-serif;">
	   <div style="width: 450px;">
		  <div style="background-color: #2196F3; width: 410px; color: white; padding: 20px;">
			 <h2 style="font-weight: lighter;">Todo </h2>
				`+todoTemplete+`
		  </div>
		  <div 
                style=" background-color: white; width: 408px; color: black; padding: 20px; font-size: smaller; border: #2196F3 1px solid;"> <br> On %s
			 </div>
	   </div>
	</body>
 </html>`,
		time.Now().String(),
	)
}
