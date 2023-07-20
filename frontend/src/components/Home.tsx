import { useState } from "react";

import MainContent from "./MainContent";
import { Player } from "./Player";
import Sidebar from "./Sidebar";
import SidebarToggle from "./SidebarToggle";

interface HomeProps {
    token: string;
}

const Home: React.FC<HomeProps> = ({ token }) => {
    const [sidebarCollapsed, setSidebarCollapsed] = useState(false);

    const toggleSidebar = () => {
        setSidebarCollapsed(!sidebarCollapsed);
    };

    return (
        <div className="container">
            <div className="row top-row">
                <div className="column column-1">
                    <SidebarToggle collapsed={sidebarCollapsed} onClick={toggleSidebar} />
                    <Sidebar collapsed={sidebarCollapsed} />
                </div>
                <div className="column column-2">
                    <MainContent />
                </div>
            </div>
            <div className="row bottom-row">
                <div className="column column-3">
                    <Player />
                </div>
            </div>
        </div>
    );
};

export default Home;