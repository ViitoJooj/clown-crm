use dioxus::prelude::*;
use super::Icon;

#[derive(Clone, PartialEq)]
pub enum FilterType {
    Text,
    Select,
    DateRange,
    Tags,
}

#[derive(Clone, PartialEq)]
pub struct FilterField {
    pub id: String,
    pub label: String,
    pub filter_type: FilterType,
    pub options: Vec<String>,
    pub value: String,
}

#[derive(Clone, PartialEq)]
pub struct ActiveFilter {
    pub id: String,
    pub label: String,
    pub value: String,
}

#[component]
pub fn FilterPanel(
    filters: Vec<FilterField>,
    active_filters: Vec<ActiveFilter>,
    on_filter_change: EventHandler<(String, String)>,
    on_apply: EventHandler<()>,
    on_reset: EventHandler<()>,
    on_remove_filter: EventHandler<String>,
    #[props(default = false)] collapsed: bool,
    on_toggle: EventHandler<()>,
) -> Element {
    rsx! {
        div {
            style: "
                background: rgba(26, 26, 26, 0.6);
                backdrop-filter: blur(20px);
                border: 1px solid rgba(255, 255, 255, 0.05);
                border-radius: 12px;
                overflow: hidden;
                transition: all 0.3s;
            ",
            
            // Header
            div {
                style: "
                    padding: 16px;
                    border-bottom: 1px solid rgba(58, 58, 58, 0.5);
                    display: flex;
                    align-items: center;
                    justify-content: space-between;
                    cursor: pointer;
                ",
                onclick: move |_| on_toggle.call(()),
                
                div {
                    style: "display: flex; align-items: center; gap: 8px;",
                    Icon {
                        name: "filter".to_string(),
                        size: "20".to_string(),
                        color: "#8B0E0E".to_string(),
                    }
                    h3 {
                        style: "
                            margin: 0;
                            font-size: 16px;
                            font-weight: 600;
                            color: white;
                        ",
                        "Filters"
                    }
                    if !active_filters.is_empty() {
                        span {
                            style: "
                                background: rgba(139, 14, 14, 0.3);
                                color: #C91A1A;
                                border: 1px solid rgba(139, 14, 14, 0.5);
                                border-radius: 999px;
                                padding: 2px 8px;
                                font-size: 12px;
                                font-weight: 600;
                            ",
                            "{active_filters.len()}"
                        }
                    }
                }
                
                Icon {
                    name: if collapsed { "chevron-down".to_string() } else { "chevron-up".to_string() },
                    size: "20".to_string(),
                    color: "#666".to_string(),
                }
            }
            
            if !collapsed {
                div {
                    style: "padding: 16px;",
                    
                    // Active filters
                    if !active_filters.is_empty() {
                        div {
                            style: "margin-bottom: 16px;",
                            
                            div {
                                style: "
                                    font-size: 12px;
                                    font-weight: 600;
                                    color: #B0B0B0;
                                    margin-bottom: 8px;
                                    text-transform: uppercase;
                                    letter-spacing: 0.5px;
                                ",
                                "Active Filters"
                            }
                            
                            div {
                                style: "display: flex; flex-wrap: wrap; gap: 8px;",
                                
                                for filter in active_filters.iter() {
                                    div {
                                        key: "{filter.id}",
                                        style: "
                                            display: inline-flex;
                                            align-items: center;
                                            gap: 6px;
                                            background: rgba(139, 14, 14, 0.2);
                                            border: 1px solid rgba(139, 14, 14, 0.4);
                                            border-radius: 6px;
                                            padding: 6px 10px;
                                            font-size: 13px;
                                        ",
                                        
                                        span {
                                            style: "color: #B0B0B0;",
                                            "{filter.label}:"
                                        }
                                        span {
                                            style: "color: white; font-weight: 600;",
                                            "{filter.value}"
                                        }
                                        
                                        button {
                                            onclick: move |_| on_remove_filter.call(filter.id.clone()),
                                            style: "
                                                background: transparent;
                                                border: none;
                                                cursor: pointer;
                                                padding: 0;
                                                display: flex;
                                                align-items: center;
                                                color: #999;
                                                transition: color 0.2s;
                                                &:hover {{ color: #C91A1A; }}
                                            ",
                                            Icon {
                                                name: "x".to_string(),
                                                size: "14".to_string(),
                                                color: "currentColor".to_string(),
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    }
                    
                    // Filter fields
                    div {
                        style: "display: flex; flex-direction: column; gap: 16px;",
                        
                        for filter in filters.iter() {
                            div {
                                key: "{filter.id}",
                                
                                label {
                                    style: "
                                        display: block;
                                        font-size: 13px;
                                        font-weight: 600;
                                        color: #B0B0B0;
                                        margin-bottom: 6px;
                                    ",
                                    "{filter.label}"
                                }
                                
                                match &filter.filter_type {
                                    FilterType::Text => rsx! {
                                        input {
                                            r#type: "text",
                                            value: "{filter.value}",
                                            oninput: move |e| on_filter_change.call((filter.id.clone(), e.value().to_string())),
                                            style: "
                                                width: 100%;
                                                padding: 10px 12px;
                                                background: rgba(42, 42, 42, 0.7);
                                                border: 1px solid rgba(58, 58, 58, 0.5);
                                                border-radius: 6px;
                                                color: white;
                                                font-size: 14px;
                                                outline: none;
                                                transition: all 0.2s;
                                                &:focus {{
                                                    border-color: rgba(139, 14, 14, 0.5);
                                                }}
                                            ",
                                        }
                                    },
                                    FilterType::Select => rsx! {
                                        select {
                                            value: "{filter.value}",
                                            onchange: move |e| on_filter_change.call((filter.id.clone(), e.value().to_string())),
                                            style: "
                                                width: 100%;
                                                padding: 10px 12px;
                                                background: rgba(42, 42, 42, 0.7);
                                                border: 1px solid rgba(58, 58, 58, 0.5);
                                                border-radius: 6px;
                                                color: white;
                                                font-size: 14px;
                                                cursor: pointer;
                                                outline: none;
                                            ",
                                            
                                            option { value: "", "-- Select --" }
                                            
                                            for opt in filter.options.iter() {
                                                option {
                                                    key: "{opt}",
                                                    value: "{opt}",
                                                    "{opt}"
                                                }
                                            }
                                        }
                                    },
                                    FilterType::DateRange => rsx! {
                                        div {
                                            style: "display: grid; grid-template-columns: 1fr 1fr; gap: 8px;",
                                            input {
                                                r#type: "date",
                                                style: "
                                                    padding: 10px 12px;
                                                    background: rgba(42, 42, 42, 0.7);
                                                    border: 1px solid rgba(58, 58, 58, 0.5);
                                                    border-radius: 6px;
                                                    color: white;
                                                    font-size: 14px;
                                                ",
                                            }
                                            input {
                                                r#type: "date",
                                                style: "
                                                    padding: 10px 12px;
                                                    background: rgba(42, 42, 42, 0.7);
                                                    border: 1px solid rgba(58, 58, 58, 0.5);
                                                    border-radius: 6px;
                                                    color: white;
                                                    font-size: 14px;
                                                ",
                                            }
                                        }
                                    },
                                    FilterType::Tags => rsx! {
                                        div {
                                            style: "display: flex; flex-wrap: wrap; gap: 8px;",
                                            
                                            for tag in filter.options.iter() {
                                                button {
                                                    key: "{tag}",
                                                    onclick: move |_| on_filter_change.call((filter.id.clone(), tag.clone())),
                                                    style: "
                                                        padding: 6px 12px;
                                                        background: {if filter.value.contains(tag) { \"rgba(139, 14, 14, 0.3)\" } else { \"rgba(42, 42, 42, 0.7)\" }};
                                                        border: 1px solid {if filter.value.contains(tag) { \"rgba(139, 14, 14, 0.5)\" } else { \"rgba(58, 58, 58, 0.5)\" }};
                                                        border-radius: 6px;
                                                        color: white;
                                                        font-size: 13px;
                                                        cursor: pointer;
                                                        transition: all 0.2s;
                                                    ",
                                                    "{tag}"
                                                }
                                            }
                                        }
                                    },
                                }
                            }
                        }
                    }
                    
                    // Action buttons
                    div {
                        style: "
                            display: flex;
                            gap: 8px;
                            margin-top: 20px;
                            padding-top: 16px;
                            border-top: 1px solid rgba(58, 58, 58, 0.3);
                        ",
                        
                        button {
                            onclick: move |_| on_apply.call(()),
                            style: "
                                flex: 1;
                                padding: 10px;
                                background: linear-gradient(135deg, #8B0E0E 0%, #C91A1A 100%);
                                border: 1px solid rgba(139, 14, 14, 0.5);
                                border-radius: 8px;
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
                            "Apply Filters"
                        }
                        
                        button {
                            onclick: move |_| on_reset.call(()),
                            style: "
                                padding: 10px 16px;
                                background: rgba(42, 42, 42, 0.7);
                                border: 1px solid rgba(58, 58, 58, 0.5);
                                border-radius: 8px;
                                color: white;
                                font-size: 14px;
                                font-weight: 600;
                                cursor: pointer;
                                transition: all 0.2s;
                            ",
                            "Reset"
                        }
                    }
                }
            }
        }
    }
}
