use dioxus::prelude::*;
use super::Icon;

#[derive(Clone, PartialEq)]
pub struct FilterOption {
    pub value: String,
    pub label: String,
}

#[component]
pub fn SearchBar(
    #[props(default = "".to_string())] value: String,
    #[props(default = "Search...".to_string())] placeholder: String,
    #[props(default = vec![])] filter_options: Vec<FilterOption>,
    #[props(default = "".to_string())] selected_filter: String,
    #[props(default = vec![])] suggestions: Vec<String>,
    #[props(default = false)] show_suggestions: bool,
    on_input: EventHandler<String>,
    on_clear: EventHandler<()>,
    on_filter_change: EventHandler<String>,
    on_suggestion_select: EventHandler<String>,
    on_search: EventHandler<()>,
) -> Element {
    rsx! {
        div {
            style: "position: relative; width: 100%;",
            
            // Main search container
            div {
                style: "
                    display: flex;
                    gap: 8px;
                    align-items: center;
                    background: rgba(42, 42, 42, 0.7);
                    backdrop-filter: blur(10px);
                    border: 1px solid rgba(58, 58, 58, 0.5);
                    border-radius: 10px;
                    padding: 4px 4px 4px 12px;
                    transition: all 0.3s;
                    &:focus-within {{
                        border-color: rgba(139, 14, 14, 0.5);
                        box-shadow: 0 0 0 3px rgba(139, 14, 14, 0.1);
                    }}
                ",
                
                // Search icon
                Icon {
                    name: "search".to_string(),
                    size: "20".to_string(),
                    color: "#666".to_string(),
                }
                
                // Filter dropdown (if options provided)
                if !filter_options.is_empty() {
                    select {
                        value: "{selected_filter}",
                        onchange: move |e| on_filter_change.call(e.value().to_string()),
                        style: "
                            background: rgba(58, 58, 58, 0.5);
                            border: 1px solid rgba(58, 58, 58, 0.5);
                            border-radius: 6px;
                            padding: 6px 12px;
                            color: white;
                            font-size: 14px;
                            cursor: pointer;
                            outline: none;
                        ",
                        
                        option { value: "", "All" }
                        
                        for opt in filter_options.iter() {
                            option {
                                key: "{opt.value}",
                                value: "{opt.value}",
                                "{opt.label}"
                            }
                        }
                    }
                }
                
                // Input field
                input {
                    r#type: "text",
                    value: "{value}",
                    placeholder: "{placeholder}",
                    oninput: move |e| on_input.call(e.value().to_string()),
                    onkeydown: move |e| {
                        if e.key() == Key::Enter {
                            on_search.call(());
                        }
                    },
                    style: "
                        flex: 1;
                        background: transparent;
                        border: none;
                        outline: none;
                        color: white;
                        font-size: 14px;
                        padding: 8px;
                    ",
                }
                
                // Clear button (if value exists)
                if !value.is_empty() {
                    button {
                        onclick: move |_| on_clear.call(()),
                        style: "
                            background: transparent;
                            border: none;
                            cursor: pointer;
                            padding: 4px;
                            display: flex;
                            align-items: center;
                            color: #666;
                            transition: color 0.2s;
                            &:hover {{ color: #999; }}
                        ",
                        Icon {
                            name: "x".to_string(),
                            size: "18".to_string(),
                            color: "currentColor".to_string(),
                        }
                    }
                }
                
                // Search button
                button {
                    onclick: move |_| on_search.call(()),
                    style: "
                        background: linear-gradient(135deg, #8B0E0E 0%, #C91A1A 100%);
                        border: 1px solid rgba(139, 14, 14, 0.5);
                        border-radius: 6px;
                        padding: 8px 16px;
                        color: white;
                        font-size: 14px;
                        font-weight: 600;
                        cursor: pointer;
                        transition: all 0.2s;
                        &:hover {{
                            transform: translateY(-1px);
                            box-shadow: 0 4px 12px rgba(139, 14, 14, 0.3);
                        }}
                    ",
                    "Search"
                }
            }
            
            // Autocomplete suggestions dropdown
            if show_suggestions && !suggestions.is_empty() {
                div {
                    style: "
                        position: absolute;
                        top: calc(100% + 4px);
                        left: 0;
                        right: 0;
                        z-index: 100;
                        background: rgba(26, 26, 26, 0.95);
                        backdrop-filter: blur(20px);
                        border: 1px solid rgba(58, 58, 58, 0.5);
                        border-radius: 8px;
                        overflow: hidden;
                        box-shadow: 0 8px 24px rgba(0, 0, 0, 0.4);
                        animation: slideDown 0.2s ease;
                    ",
                    
                    for (idx, suggestion) in suggestions.iter().enumerate() {
                        div {
                            key: "{idx}",
                            onclick: move |_| on_suggestion_select.call(suggestion.clone()),
                            style: "
                                padding: 12px 16px;
                                color: white;
                                cursor: pointer;
                                transition: background 0.2s;
                                &:hover {{
                                    background: rgba(139, 14, 14, 0.2);
                                }}
                                {if idx > 0 { \"border-top: 1px solid rgba(58, 58, 58, 0.3);\" } else { \"\" }}
                            ",
                            
                            div {
                                style: "display: flex; align-items: center; gap: 8px;",
                                Icon {
                                    name: "search".to_string(),
                                    size: "16".to_string(),
                                    color: "#666".to_string(),
                                }
                                span { "{suggestion}" }
                            }
                        }
                    }
                }
            }
        }
    }
}

#[component]
pub fn CompactSearchBar(
    #[props(default = "".to_string())] value: String,
    #[props(default = "Search...".to_string())] placeholder: String,
    on_input: EventHandler<String>,
    on_clear: EventHandler<()>,
) -> Element {
    rsx! {
        div {
            style: "position: relative; width: 100%;",
            
            input {
                r#type: "text",
                value: "{value}",
                placeholder: "{placeholder}",
                oninput: move |e| on_input.call(e.value().to_string()),
                style: "
                    width: 100%;
                    padding: 10px 40px 10px 40px;
                    border-radius: 8px;
                    border: 1px solid rgba(58, 58, 58, 0.5);
                    background: rgba(42, 42, 42, 0.7);
                    color: white;
                    font-size: 14px;
                    outline: none;
                    transition: all 0.3s;
                    &:focus {{
                        border-color: rgba(139, 14, 14, 0.5);
                        box-shadow: 0 0 0 3px rgba(139, 14, 14, 0.1);
                    }}
                ",
            }
            
            div {
                style: "
                    position: absolute;
                    left: 12px;
                    top: 50%;
                    transform: translateY(-50%);
                    pointer-events: none;
                ",
                Icon {
                    name: "search".to_string(),
                    size: "18".to_string(),
                    color: "#666".to_string(),
                }
            }
            
            if !value.is_empty() {
                button {
                    onclick: move |_| on_clear.call(()),
                    style: "
                        position: absolute;
                        right: 12px;
                        top: 50%;
                        transform: translateY(-50%);
                        background: transparent;
                        border: none;
                        cursor: pointer;
                        padding: 4px;
                        display: flex;
                        align-items: center;
                        color: #666;
                        transition: color 0.2s;
                        &:hover {{ color: #999; }}
                    ",
                    Icon {
                        name: "x".to_string(),
                        size: "18".to_string(),
                        color: "currentColor".to_string(),
                    }
                }
            }
        }
    }
}
