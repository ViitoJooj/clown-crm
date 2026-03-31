#[derive(Clone, Routable, Debug, PartialEq)]
#[rustfmt::skip]
pub enum Route {
    #[layout(MainLayout)]
        #[route("/")]
        Dashboard {},
        #[route("/chats")]
        Chats {},
        #[route("/admin")]
        Admin {},
        #[route("/profile")]
        Profile {},
    #[route("/login")]
    LoginPage {},
    #[route("/register")]
    RegisterPage {},
    #[route("/:..route")]
    PageNotFound { route: Vec<String> },
}

#[component]
fn MainLayout() -> Element {
    rsx! {
        // Main app layout will be rendered here
        Outlet::<Route> {}
    }
}

#[component]
fn Dashboard() -> Element {
    rsx! {
        div {
            "Dashboard"
        }
    }
}

#[component]
fn Chats() -> Element {
    rsx! {
        div {
            "Chats"
        }
    }
}

#[component]
fn Admin() -> Element {
    rsx! {
        div {
            "Admin Panel"
        }
    }
}

#[component]
fn Profile() -> Element {
    rsx! {
        div {
            "User Profile"
        }
    }
}

#[component]
fn PageNotFound(route: Vec<String>) -> Element {
    rsx! {
        div {
            h1 { "Page Not Found" }
            p { "The page {route.join(\"/\")} does not exist." }
        }
    }
}
