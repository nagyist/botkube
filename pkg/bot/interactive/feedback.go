package interactive

// Feedback generates Message structure.
func Feedback(botName string) Message {
	btnBuilder := buttonBuilder{botName: botName}
	return Message{
		Sections: []Section{
			{
				Base: Base{
					Body: Body{
						Plaintext: ":wave: Hey, what's your experience with Botkube so far?",
					},
				},
				Buttons: []Button{
					btnBuilder.ForURL("Give feedback", "https://feedback.botkube.io", ButtonStylePrimary),
				},
			},
		},
		OnlyVisibleForYou: true,
	}
}