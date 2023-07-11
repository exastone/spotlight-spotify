import { useState } from "react";
import { invoke } from "@tauri-apps/api/tauri";

import "./App.css";
import { Player } from "./components/Player";
import Sidebar from "./components/Sidebar";
import SidebarToggle from "./components/SidebarToggle";
import MainContent from "./components/MainContent";

// https://developer.spotify.com/documentation/web-api/howtos/web-app-profile

function App() {

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
}

export default App;
