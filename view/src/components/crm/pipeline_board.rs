use dioxus::prelude::*;
use crate::components::common::Icon;

#[derive(Clone, PartialEq)]
pub struct Deal {
    pub id: String,
    pub title: String,
    pub value: f64,
    pub company: String,
}

#[derive(Clone, PartialEq)]
pub struct PipelineStage {
    pub id: String,
    pub name: String,
    pub deals: Vec<Deal>,
    pub color: String,
}

#[derive(Props, Clone, PartialEq)]
pub struct PipelineBoardProps {
    pub stages: Vec<PipelineStage>,
    #[props(default = None)]
    pub on_deal_click: Option<EventHandler<String>>,
    #[props(default = None)]
    pub on_deal_move: Option<EventHandler<(String, String)>>,
}

#[component]
pub fn PipelineBoard(props: PipelineBoardProps) -> Element {
    rsx! {
        div {
            class: "pipeline-board",
            style: r#"
                display: flex;
                gap: 20px;
                padding: 20px;
                overflow-x: auto;
                min-height: 600px;
            "#,

            for stage in props.stages.iter() {
                div {
                    key: "{stage.id}",
                    class: "pipeline-stage",
                    style: r#"
                        min-width: 300px;
                        max-width: 300px;
                        background: linear-gradient(135deg, rgba(26, 26, 26, 0.95) 0%, rgba(42, 42, 42, 0.85) 100%);
                        backdrop-filter: blur(10px);
                        -webkit-backdrop-filter: blur(10px);
                        border: 1px solid rgba(58, 58, 58, 0.5);
                        border-radius: 16px;
                        padding: 16px;
                        display: flex;
                        flex-direction: column;
                        transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
                    "#,

                    div {
                        style: "display: flex; justify-content: space-between; align-items: center; margin-bottom: 16px; padding-bottom: 12px; border-bottom: 1px solid rgba(58, 58, 58, 0.5);",
                        
                        div {
                            style: "display: flex; align-items: center; gap: 8px;",
                            
                            div {
                                style: format!("width: 12px; height: 12px; border-radius: 50%; background: {};", stage.color),
                            }

                            h3 {
                                style: "font-size: 15px; font-weight: 700; color: #FFFFFF;",
                                "{stage.name}"
                            }
                        }

                        span {
                            style: format!(r#"
                                display: inline-flex;
                                align-items: center;
                                justify-content: center;
                                min-width: 24px;
                                height: 24px;
                                padding: 0 8px;
                                border-radius: 999px;
                                font-size: 12px;
                                font-weight: 700;
                                background: {};
                                color: #FFFFFF;
                            "#, stage.color),
                            "{stage.deals.len()}"
                        }
                    }

                    div {
                        class: "pipeline-stage-deals",
                        style: r#"
                            flex: 1;
                            display: flex;
                            flex-direction: column;
                            gap: 12px;
                            overflow-y: auto;
                            padding: 4px;
                        "#,

                        if stage.deals.is_empty() {
                            div {
                                style: r#"
                                    display: flex;
                                    flex-direction: column;
                                    align-items: center;
                                    justify-content: center;
                                    padding: 40px 20px;
                                    color: #707070;
                                    text-align: center;
                                "#,
                                Icon { name: "briefcase".to_string(), size: "32".to_string(), color: "#707070".to_string(), class: "".to_string() }
                                p { style: "margin-top: 12px; font-size: 13px;", "No deals in this stage" }
                            }
                        } else {
                            for deal in stage.deals.iter() {
                                div {
                                    key: "{deal.id}",
                                    class: "pipeline-deal-card",
                                    onclick: move |_| {
                                        if let Some(handler) = &props.on_deal_click {
                                            handler.call(deal.id.clone());
                                        }
                                    },
                                    style: r#"
                                        background: rgba(42, 42, 42, 0.7);
                                        border: 1px solid rgba(58, 58, 58, 0.5);
                                        border-radius: 12px;
                                        padding: 16px;
                                        cursor: grab;
                                        transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
                                    "#,

                                    h4 {
                                        style: "font-size: 14px; font-weight: 600; color: #FFFFFF; margin-bottom: 8px;",
                                        "{deal.title}"
                                    }

                                    if !deal.company.is_empty() {
                                        p {
                                            style: "font-size: 12px; color: #B0B0B0; margin-bottom: 12px; display: flex; align-items: center; gap: 6px;",
                                            Icon { name: "building".to_string(), size: "12".to_string(), color: "#B0B0B0".to_string(), class: "".to_string() }
                                            "{deal.company}"
                                        }
                                    }

                                    div {
                                        style: "display: flex; justify-content: space-between; align-items: center; padding-top: 12px; border-top: 1px solid rgba(58, 58, 58, 0.5);",
                                        
                                        div {
                                            style: "display: flex; align-items: center; gap: 6px;",
                                            Icon { name: "dollar-sign".to_string(), size: "16".to_string(), color: "#10B981".to_string(), class: "".to_string() }
                                            span {
                                                style: "font-size: 16px; font-weight: 700; color: #10B981;",
                                                "${deal.value:.0}"
                                            }
                                        }

                                        Icon { name: "more-vertical".to_string(), size: "16".to_string(), color: "#707070".to_string(), class: "".to_string() }
                                    }
                                }
                            }
                        }
                    }

                    div {
                        style: r#"
                            margin-top: 12px;
                            padding-top: 12px;
                            border-top: 1px solid rgba(58, 58, 58, 0.5);
                            text-align: center;
                        "#,
                        
                        button {
                            style: r#"
                                width: 100%;
                                padding: 10px;
                                border-radius: 10px;
                                background: rgba(139, 14, 14, 0.1);
                                border: 1px dashed rgba(139, 14, 14, 0.5);
                                color: #C91A1A;
                                display: flex;
                                align-items: center;
                                justify-content: center;
                                gap: 6px;
                                font-size: 13px;
                                font-weight: 600;
                                cursor: pointer;
                                transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
                            "#,
                            Icon { name: "plus".to_string(), size: "16".to_string(), color: "currentColor".to_string(), class: "".to_string() }
                            "Add Deal"
                        }
                    }
                }
            }

            style {
                r#"
                .pipeline-board::-webkit-scrollbar {
                    height: 8px;
                }
                
                .pipeline-stage:hover {
                    border-color: rgba(139, 14, 14, 0.3);
                    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
                }
                
                .pipeline-deal-card:hover {
                    transform: translateY(-2px);
                    box-shadow: 0 6px 20px rgba(0, 0, 0, 0.4);
                    border-color: rgba(139, 14, 14, 0.4);
                }
                
                .pipeline-deal-card:active {
                    cursor: grabbing;
                    transform: scale(1.02);
                }
                
                .pipeline-stage button:hover {
                    background: rgba(139, 14, 14, 0.2);
                    border-color: rgba(139, 14, 14, 0.7);
                    transform: translateY(-2px);
                }
                "#
            }
        }
    }
}
