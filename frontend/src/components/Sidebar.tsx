import { HomeIcon } from "./HomeIcon";
import { LibraryIcon } from "./LibraryIcon";
import { SearchIcon } from "./SearchIcon";

import "../styles/sidebar.css";
import ListItem from "./ListItem";

interface SidebarProps {
    collapsed: boolean;
}

const Sidebar: React.FC<SidebarProps> = ({ collapsed }) => {

    return (
        <div className={`sidebar ${collapsed ? "collapsed" : ""}`}>
            <div className="logo-spotlight">
                <p>Spotlight</p>
            </div>

            <div className="menu">
                <div className="menu-item home">
                    <button className="button home">
                        <HomeIcon />
                        Home
                    </button>
                </div>
                <div className="menu-item search">
                    <button className="button search">
                        <SearchIcon />
                        Search
                    </button>
                </div>
            </div>

            <div className="menu user-library">
                <div className="menu-item library">
                    <button className="button library">
                        <LibraryIcon />
                        Library
                    </button>
                </div>
                <ul className="library-list">
                    <ListItem
                        image={"src/assets/TLB-hit.jpeg"}
                        title={"1: *(char*)0 = 0"}
                        artist={"TLB Hit💥"}
                        type={"Podcast"}
                    />
                    <ListItem
                        image={"src/assets/liked-songs-64.png"}
                        title={"Liked Songs ♥️ "}
                        artist={"By you"}
                        type={"Playlist"}
                    />
                    <ListItem
                        image={"src/assets/protea.png"}
                        title={"Protea"}
                        artist={"Kota the Friend"}
                        type={"Album"}
                    />
                    <ListItem
                        image={"src/assets/sunburn.png"}
                        title={"Sunburn"}
                        artist={"Dominic Fike"}
                        type={"Album"}
                    />
                    <ListItem
                        image={"src/assets/rustacean-station.png"}
                        title={"Rustacean Station"}
                        artist={"Rustacean Station"}
                        type={"Podcast"}
                    />
                    <ListItem
                        image={"src/assets/ilovelife-thankyou.webp"}
                        title={"I Love Life, Thank You"}
                        artist={"Mac Miller"}
                        type={"Album"}
                    />

                </ul>
            </div>
        </div>
    );
};

export default Sidebar;