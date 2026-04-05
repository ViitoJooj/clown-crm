// Exemplo de uso do sistema de temas e novos componentes

use crate::components::common::Icon;
use crate::components::*;
use crate::styles::theme::ThemeVariant;
use dioxus::prelude::*;

/// Página de demonstração dos componentes e temas
#[component]
pub fn ComponentShowcase() -> Element {
    let mut current_theme = use_signal(|| ThemeVariant::Burgundy);

    rsx! {
        div {
            style: "min-height: 100vh; padding: 40px; background: linear-gradient(180deg, #0A0A0A 0%, #1A0505 100%);",

            // Header with theme selector
            div {
                style: "display: flex; justify-content: space-between; align-items: center; margin-bottom: 40px;",
                h1 {
                    style: "font-size: 36px; font-weight: 900; background: linear-gradient(135deg, #8B0E0E, #C91A1A); -webkit-background-clip: text; -webkit-text-fill-color: transparent;",
                    "Component Showcase"
                }
                ThemeSelector {
                    current_theme
                }
            }

            // Buttons Section
            Card {
                glass: false,
                div {
                    style: "margin-bottom: 30px;",
                    h2 {
                        style: "margin-bottom: 20px; color: #B0B0B0;",
                        "Buttons"
                    }
                    div {
                        style: "display: flex; flex-wrap: wrap; gap: 12px;",
                        Button {
                            class: "primary".to_string(),
                            onclick: |_| {},
                            "Primary Button"
                        }
                        Button {
                            class: "secondary".to_string(),
                            onclick: |_| {},
                            "Secondary Button"
                        }
                        Button {
                            class: "danger".to_string(),
                            onclick: |_| {},
                            "Danger Button"
                        }
                        Button {
                            class: "success".to_string(),
                            onclick: |_| {},
                            "Success Button"
                        }
                        Button {
                            class: "ghost".to_string(),
                            onclick: |_| {},
                            "Ghost Button"
                        }
                        Button {
                            class: "primary".to_string(),
                            disabled: true,
                            onclick: |_| {},
                            "Disabled"
                        }
                    }
                }
            }

            // Badges Section
            Card {
                glass: true,
                div {
                    style: "margin: 30px 0;",
                    h2 {
                        style: "margin-bottom: 20px; color: #B0B0B0;",
                        "Badges"
                    }
                    div {
                        style: "display: flex; flex-wrap: wrap; gap: 10px;",
                        Badge {
                            variant: "success".to_string(),
                            div {
                                style: "display: flex; align-items: center; gap: 4px;",
                                Icon { name: "check".to_string(), size: "14".to_string() }
                                "Success"
                            }
                        }
                        Badge {
                            variant: "error".to_string(),
                            div {
                                style: "display: flex; align-items: center; gap: 4px;",
                                Icon { name: "x".to_string(), size: "14".to_string() }
                                "Error"
                            }
                        }
                        Badge {
                            variant: "warning".to_string(),
                            div {
                                style: "display: flex; align-items: center; gap: 4px;",
                                Icon { name: "alert-triangle".to_string(), size: "14".to_string() }
                                "Warning"
                            }
                        }
                        Badge {
                            variant: "info".to_string(),
                            div {
                                style: "display: flex; align-items: center; gap: 4px;",
                                Icon { name: "info".to_string(), size: "14".to_string() }
                                "Info"
                            }
                        }
                        Badge {
                            variant: "primary".to_string(),
                            div {
                                style: "display: flex; align-items: center; gap: 4px;",
                                Icon { name: "star".to_string(), size: "14".to_string() }
                                "Primary"
                            }
                        }
                        Badge {
                            variant: "default".to_string(),
                            "Default"
                        }
                    }
                }
            }

            // Inputs Section
            Card {
                glass: false,
                div {
                    style: "margin: 30px 0;",
                    h2 {
                        style: "margin-bottom: 20px; color: #B0B0B0;",
                        "Inputs"
                    }
                    div {
                        style: "display: flex; flex-direction: column; gap: 16px; max-width: 400px;",
                        Input {
                            input_type: "text".to_string(),
                            placeholder: "Enter your name".to_string(),
                            value: "".to_string(),
                            onchange: |_| {}
                        }
                        Input {
                            input_type: "email".to_string(),
                            placeholder: "your@email.com".to_string(),
                            value: "".to_string(),
                            onchange: |_| {}
                        }
                        Input {
                            input_type: "password".to_string(),
                            placeholder: "Password".to_string(),
                            value: "".to_string(),
                            onchange: |_| {}
                        }
                    }
                }
            }

            // Avatars Section
            Card {
                glass: true,
                div {
                    style: "margin: 30px 0;",
                    h2 {
                        style: "margin-bottom: 20px; color: #B0B0B0;",
                        "Avatars"
                    }
                    div {
                        style: "display: flex; align-items: center; gap: 16px;",
                        Avatar {
                            src: "".to_string(),
                            alt: "John Doe".to_string(),
                            size: "40px".to_string()
                        }
                        Avatar {
                            src: "".to_string(),
                            alt: "Jane Smith".to_string(),
                            size: "50px".to_string()
                        }
                        Avatar {
                            src: "".to_string(),
                            alt: "Bob Wilson".to_string(),
                            size: "60px".to_string()
                        }
                    }
                }
            }

            // Loading Spinner
            Card {
                glass: false,
                div {
                    style: "margin: 30px 0;",
                    h2 {
                        style: "margin-bottom: 20px; color: #B0B0B0;",
                        "Loading Spinners"
                    }
                    div {
                        style: "display: flex; align-items: center; gap: 24px;",
                        LoadingSpinner { size: "24px".to_string() }
                        LoadingSpinner { size: "32px".to_string() }
                        LoadingSpinner { size: "48px".to_string() }
                    }
                }
            }

            // Dividers
            Card {
                glass: true,
                div {
                    style: "margin: 30px 0;",
                    h2 {
                        style: "margin-bottom: 20px; color: #B0B0B0;",
                        "Dividers"
                    }
                    p { "Content above divider" }
                    Divider {}
                    p { "Content below divider" }
                }
            }
        }
    }
}
