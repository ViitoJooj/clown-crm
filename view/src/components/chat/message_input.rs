use dioxus::prelude::*;

#[component]
pub fn MessageInput(
    onsubmit: EventHandler<String>,
) -> Element {
    let mut message = use_signal(|| "".to_string());

    let handle_send = move |_: MouseEvent| {
        if !message.read().trim().is_empty() {
            onsubmit.call(message.read().trim().to_string());
            message.set("".to_string());
        }
    };

    let container_style = r#"
        display: flex;
        align-items: flex-end;
        gap: 12px;
        padding: 12px 16px;
        background-color: #1A1A1A;
        border-top: 1px solid #2A2A2A;
    "#;

    let input_container_style = r#"
        flex: 1;
        display: flex;
        align-items: center;
        gap: 8px;
        background-color: #2A2A2A;
        border-radius: 24px;
        padding: 8px 16px;
    "#;

    let input_style = r#"
        flex: 1;
        background: none;
        border: none;
        color: #FFFFFF;
        font-size: 14px;
        padding: 8px 0;
        outline: none;
        font-family: inherit;
    "#;

    let send_button_style = r#"
        background: none;
        border: none;
        color: #8B0E0E;
        cursor: pointer;
        font-size: 18px;
        padding: 8px;
        border-radius: 50%;
        transition: all 0.2s ease;
        display: flex;
        align-items: center;
        justify-content: center;
    "#;

    rsx! {
        div {
            style: container_style,
            
            // Attachment button
            button {
                style: "background: none; border: none; color: #8B0E0E; cursor: pointer; font-size: 18px; padding: 8px; border-radius: 50%;",
                "📎"
            }

            // Input area
            div {
                style: input_container_style,
                input {
                    r#type: "text",
                    placeholder: "Message...",
                    value: message.read().to_string(),
                    onchange: move |e: FormEvent| message.set(e.value()),
                    onkeydown: move |e: KeyboardEvent| {
                        if e.key().to_string() == "Enter" {
                            if !message.read().trim().is_empty() {
                                onsubmit.call(message.read().trim().to_string());
                                message.set("".to_string());
                            }
                        }
                    },
                    style: input_style,
                }
                
                // Emoji button
                button {
                    style: "background: none; border: none; color: #8B0E0E; cursor: pointer; font-size: 16px; padding: 4px;",
                    "😊"
                }
            }

            // Send button
            button {
                style: send_button_style,
                onclick: handle_send,
                "📤"
            }
        }
    }
}

