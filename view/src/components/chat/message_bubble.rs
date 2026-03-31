use dioxus::prelude::*;

#[component]
pub fn MessageBubble(
    #[props(default = "".to_string())] message: String,
    #[props(default = false)] is_sent: bool,
    #[props(default = "".to_string())] time: String,
) -> Element {
    let bubble_style = if is_sent {
        r#"
        display: flex;
        justify-content: flex-end;
        margin-bottom: 8px;
        "#
    } else {
        r#"
        display: flex;
        justify-content: flex-start;
        margin-bottom: 8px;
        "#
    };

    let content_style = if is_sent {
        r#"
        max-width: 70%;
        background-color: #8B0E0E;
        color: #FFFFFF;
        padding: 8px 12px;
        border-radius: 18px;
        border-bottom-right-radius: 4px;
        word-wrap: break-word;
        "#
    } else {
        r#"
        max-width: 70%;
        background-color: #2A2A2A;
        color: #FFFFFF;
        padding: 8px 12px;
        border-radius: 18px;
        border-bottom-left-radius: 4px;
        word-wrap: break-word;
        "#
    };

    rsx! {
        div {
            style: bubble_style,
            div {
                style: content_style,
                p {
                    style: "margin: 0 0 4px 0; font-size: 15px; line-height: 1.4;",
                    "{message}"
                }
                span {
                    style: "font-size: 11px; color: #B0B0B0; display: block; text-align: right;",
                    "{time}"
                }
            }
        }
    }
}
