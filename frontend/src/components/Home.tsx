import { useState } from "react";
import { Sample } from "./Sample";

export function Home() {
    const [name, setName] = useState("");
    return (
        // <p>Welcome to Spotlight {name}</p>
        <Sample />
    );
}