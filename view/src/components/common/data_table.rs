use dioxus::prelude::*;
use super::Icon;

#[derive(Clone, PartialEq)]
pub enum SortDirection {
    Asc,
    Desc,
    None,
}

#[derive(Clone, PartialEq)]
pub struct ColumnDef {
    pub key: String,
    pub label: String,
    pub sortable: bool,
    pub width: Option<String>,
}

#[component]
pub fn DataTable(
    columns: Vec<ColumnDef>,
    #[props(default = vec![])] selected_rows: Vec<String>,
    #[props(default = "".to_string())] sort_column: String,
    #[props(default = SortDirection::None)] sort_direction: SortDirection,
    #[props(default = 1)] current_page: i32,
    #[props(default = 10)] per_page: i32,
    #[props(default = 0)] total_items: i32,
    #[props(default = "".to_string())] search_query: String,
    on_sort: EventHandler<String>,
    on_row_select: EventHandler<String>,
    on_select_all: EventHandler<()>,
    on_page_change: EventHandler<i32>,
    on_search: EventHandler<String>,
    children: Element,
) -> Element {
    let total_pages = (total_items as f32 / per_page as f32).ceil() as i32;
    let start_item = ((current_page - 1) * per_page) + 1;
    let end_item = (current_page * per_page).min(total_items);

    rsx! {
        div {
            style: "width: 100%;",
            
            // Search bar
            div {
                style: "margin-bottom: 16px;",
                div {
                    style: "position: relative; max-width: 400px;",
                    input {
                        r#type: "text",
                        placeholder: "Search...",
                        value: "{search_query}",
                        oninput: move |e| on_search.call(e.value().to_string()),
                        style: "
                            width: 100%;
                            padding: 10px 10px 10px 40px;
                            border-radius: 8px;
                            border: 1px solid rgba(58, 58, 58, 0.5);
                            background: rgba(42, 42, 42, 0.7);
                            color: white;
                            font-size: 14px;
                        ",
                    }
                    div {
                        style: "position: absolute; left: 12px; top: 50%; transform: translateY(-50%); pointer-events: none;",
                        Icon { name: "search".to_string(), size: "18".to_string(), color: "#666".to_string() }
                    }
                }
            }

            // Table container
            div {
                style: "
                    background: rgba(26, 26, 26, 0.6);
                    backdrop-filter: blur(20px);
                    border: 1px solid rgba(255, 255, 255, 0.05);
                    border-radius: 12px;
                    overflow: hidden;
                ",
                
                // Table
                div {
                    style: "overflow-x: auto;",
                    table {
                        style: "
                            width: 100%;
                            border-collapse: collapse;
                        ",
                        
                        // Header
                        thead {
                            tr {
                                style: "
                                    background: rgba(42, 42, 42, 0.5);
                                    border-bottom: 1px solid rgba(58, 58, 58, 0.5);
                                ",
                                
                                // Select all checkbox
                                th {
                                    style: "padding: 16px; text-align: left; width: 40px;",
                                    input {
                                        r#type: "checkbox",
                                        checked: selected_rows.len() > 0 && total_items > 0,
                                        onclick: move |_| on_select_all.call(()),
                                        style: "cursor: pointer; width: 16px; height: 16px;",
                                    }
                                }
                                
                                // Column headers
                                for col in columns.iter() {
                                    th {
                                        key: "{col.key}",
                                        style: "
                                            padding: 16px;
                                            text-align: left;
                                            font-weight: 600;
                                            color: #B0B0B0;
                                            white-space: nowrap;
                                            {col.width.as_ref().map(|w| format!(\"width: {};\", w)).unwrap_or_default()}
                                        ",
                                        
                                        if col.sortable {
                                            div {
                                                style: "
                                                    display: flex;
                                                    align-items: center;
                                                    gap: 8px;
                                                    cursor: pointer;
                                                    user-select: none;
                                                ",
                                                onclick: move |_| on_sort.call(col.key.clone()),
                                                
                                                span { "{col.label}" }
                                                
                                                if sort_column == col.key {
                                                    Icon {
                                                        name: if matches!(sort_direction, SortDirection::Asc) {
                                                            "chevron-up".to_string()
                                                        } else {
                                                            "chevron-down".to_string()
                                                        },
                                                        size: "16".to_string(),
                                                        color: "#8B0E0E".to_string(),
                                                    }
                                                } else {
                                                    Icon {
                                                        name: "chevron-down".to_string(),
                                                        size: "16".to_string(),
                                                        color: "#444".to_string(),
                                                    }
                                                }
                                            }
                                        } else {
                                            span { "{col.label}" }
                                        }
                                    }
                                }
                            }
                        }
                        
                        // Body (rendered by children)
                        tbody {
                            {children}
                        }
                    }
                }
                
                // Pagination footer
                div {
                    style: "
                        padding: 16px;
                        border-top: 1px solid rgba(58, 58, 58, 0.5);
                        display: flex;
                        align-items: center;
                        justify-content: space-between;
                        flex-wrap: wrap;
                        gap: 16px;
                    ",
                    
                    // Items info
                    div {
                        style: "color: #B0B0B0; font-size: 14px;",
                        if total_items > 0 {
                            "Showing {start_item}-{end_item} of {total_items} items"
                        } else {
                            "No items found"
                        }
                    }
                    
                    // Page controls
                    if total_pages > 1 {
                        div {
                            style: "display: flex; align-items: center; gap: 8px;",
                            
                            // Previous button
                            button {
                                disabled: current_page <= 1,
                                onclick: move |_| on_page_change.call(current_page - 1),
                                style: "
                                    padding: 8px 12px;
                                    background: rgba(42, 42, 42, 0.7);
                                    border: 1px solid rgba(58, 58, 58, 0.5);
                                    border-radius: 6px;
                                    color: white;
                                    cursor: pointer;
                                    transition: all 0.2s;
                                    {if current_page <= 1 { \"opacity: 0.5; cursor: not-allowed;\" } else { \"\" }}
                                ",
                                Icon { name: "chevron-left".to_string(), size: "16".to_string() }
                            }
                            
                            // Page numbers
                            for page in 1..=total_pages {
                                button {
                                    key: "{page}",
                                    onclick: move |_| on_page_change.call(page),
                                    style: "
                                        padding: 8px 12px;
                                        background: {if page == current_page { \"rgba(139, 14, 14, 0.3)\" } else { \"rgba(42, 42, 42, 0.7)\" }};
                                        border: 1px solid {if page == current_page { \"rgba(139, 14, 14, 0.5)\" } else { \"rgba(58, 58, 58, 0.5)\" }};
                                        border-radius: 6px;
                                        color: white;
                                        cursor: pointer;
                                        min-width: 36px;
                                        transition: all 0.2s;
                                    ",
                                    "{page}"
                                }
                            }
                            
                            // Next button
                            button {
                                disabled: current_page >= total_pages,
                                onclick: move |_| on_page_change.call(current_page + 1),
                                style: "
                                    padding: 8px 12px;
                                    background: rgba(42, 42, 42, 0.7);
                                    border: 1px solid rgba(58, 58, 58, 0.5);
                                    border-radius: 6px;
                                    color: white;
                                    cursor: pointer;
                                    transition: all 0.2s;
                                    {if current_page >= total_pages { \"opacity: 0.5; cursor: not-allowed;\" } else { \"\" }}
                                ",
                                Icon { name: "chevron-right".to_string(), size: "16".to_string() }
                            }
                        }
                    }
                }
            }
        }
    }
}

#[component]
pub fn DataTableRow(
    row_id: String,
    #[props(default = false)] selected: bool,
    on_select: EventHandler<String>,
    children: Element,
) -> Element {
    rsx! {
        tr {
            key: "{row_id}",
            style: "
                border-bottom: 1px solid rgba(58, 58, 58, 0.3);
                transition: background 0.2s;
                &:hover {{ background: rgba(139, 14, 14, 0.1); }}
                {if selected { \"background: rgba(139, 14, 14, 0.15);\" } else { \"\" }}
            ",
            
            // Checkbox cell
            td {
                style: "padding: 16px; width: 40px;",
                input {
                    r#type: "checkbox",
                    checked: selected,
                    onclick: move |_| on_select.call(row_id.clone()),
                    style: "cursor: pointer; width: 16px; height: 16px;",
                }
            }
            
            // Content cells
            {children}
        }
    }
}

#[component]
pub fn DataTableCell(
    #[props(default = "left".to_string())] align: String,
    children: Element,
) -> Element {
    rsx! {
        td {
            style: "
                padding: 16px;
                color: white;
                text-align: {align};
            ",
            {children}
        }
    }
}
