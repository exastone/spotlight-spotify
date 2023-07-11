import "../styles/topbar.css";

export function TopBar() {
    return (
        <div className="topbar">
            <div className="prev-next-buttons">
                <button type="button" className="chevron-left">
                    <h1> {"<"} </h1>
                </button>
                <button type="button" className="chevron-right">
                    <h1> {">"} </h1>
                </button>
            </div>

            <div className="navbar">
                <ul>
                    <li className="divider">|</li>
                </ul>
                <button type="button">Log In</button>
            </div>
        </div>
    );
}