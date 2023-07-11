import { useState } from "react";
import "../styles/sample.css";

import { AlbumCard } from "./AlbumCard";
import { Player } from "./Player";
import Sidebar from "./Sidebar";
import { TopBar } from "./TopBar";
import SidebarToggle from "./SidebarToggle";

export function Sample() {
    const [sidebarCollapsed, setSidebarCollapsed] = useState(false);

    const toggleSidebar = () => {
        setSidebarCollapsed(!sidebarCollapsed);
    };

    return (
        <div>
            <SidebarToggle collapsed={sidebarCollapsed} onClick={toggleSidebar} />

            <Sidebar collapsed={sidebarCollapsed} />
            <Player />
            <div className="main-container">
                <TopBar />

                <div className="spotify-playlists">
                    <h2>Spotify Playlists</h2>

                    <AlbumCard />

                    <div className="list">

                        <div className="item">
                            <img src="" />
                            <div className="play">
                                <span className="fa fa-play"></span>
                            </div>
                            <h4>RapCaviar</h4>
                            <p>New Music from Lil Baby, Juice WRLD an...</p>
                        </div>
                        <div className="item">
                            <img src="" />
                            <div className="play">
                                <span className="fa fa-play"></span>
                            </div>
                            <h4>Chill Hits</h4>
                            <p>Kick back to the best new and recent chill...</p>
                        </div>

                        <div className="item">
                            <img src="" />
                            <div className="play">
                                <span className="fa fa-play"></span>
                            </div>
                            <h4>Viva Latino</h4>
                            <p>Today's top Latin hits elevando nuestra...</p>
                        </div>

                        <div className="item">
                            <img src="" />
                            <div className="play">
                                <span className="fa fa-play"></span>
                            </div>
                            <h4>Mega Hit Mix</h4>
                            <p>A mega mix of 75 favorites from the last...</p>
                        </div>

                        <div className="item">
                            <img src="" />
                            <div className="play">
                                <span className="fa fa-play"></span>
                            </div>
                            <h4>All out 80s</h4>
                            <p>The biggest songs of the 1090s.</p>
                        </div>
                    </div>
                </div>

                <div className="spotify-playlists">
                    <h2>Focus</h2>
                    <div className="list">
                        <div className="item">
                            <img src="" />
                            <div className="play">
                                <span className="fa fa-play"></span>
                            </div>
                            <h4>Peaceful Piano</h4>
                            <p>Relax and indulge with beautiful piano pieces</p>
                        </div>

                        <div className="item">
                            <img src="" />
                            <div className="play">
                                <span className="fa fa-play"></span>
                            </div>
                            <h4>Deep Focus</h4>
                            <p>Keep calm and focus with ambient and pos...</p>
                        </div>



                        <div className="item">
                            <img src="" />
                            <div className="play">
                                <span className="fa fa-play"></span>
                            </div>
                            <h4>Focus Flow</h4>
                            <p>Uptempo instrumental hip hop beats.</p>
                        </div>

                        <div className="item">
                            <img src="" />
                            <div className="play">
                                <span className="fa fa-play"></span>
                            </div>
                            <h4>Calm Before The Storm</h4>
                            <p>Calm before the storm music.</p>
                        </div>

                        <div className="item">
                            <img src="" />
                            <div className="play">
                                <span className="fa fa-play"></span>
                            </div>
                            <h4>Beats to think to</h4>
                            <p>Focus with deep techno and tech house.</p>
                        </div>
                    </div>
                </div>

                <div className="spotify-playlists">
                    <h2>Mood</h2>
                    <div className="list">
                        <div className="item">
                            <img src="" />
                            <div className="play">
                                <span className="fa fa-play"></span>
                            </div>
                            <h4>Mood Booster</h4>
                            <p>Get happy with today's dose of feel-good...</p>
                        </div>

                        <div className="item">
                            <img src="" />
                            <div className="play">
                                <span className="fa fa-play"></span>
                            </div>
                            <h4>Feelin' Good</h4>
                            <p>Feel good with this positively timeless...</p>
                        </div>

                        <div className="item">
                            <img src="" />
                            <div className="play">
                                <span className="fa fa-play"></span>
                            </div>
                            <h4>Dark & Stormy</h4>
                            <p>Beautifully dark, dramatic tracks.</p>
                        </div>

                        <div className="item">
                            <img src="" />
                            <div className="play">
                                <span className="fa fa-play"></span>
                            </div>
                            <h4>Feel Good Piano</h4>
                            <p>Happy vibes for an upbeat morning.</p>
                        </div>



                        <div className="item">
                            <img src="" />
                            <div className="play">
                                <span className="fa fa-play"></span>
                            </div>
                            <h4>Feel-Good Indie Rock</h4>
                            <p>The best indie rock vibes - classic and...</p>
                        </div>

                    </div>
                    <hr />
                </div>
            </div>
        </div>
    );
}