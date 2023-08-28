import { useEffect, useState } from "react";

import "./App.css";
// import Home from "./components/Home";
import { Login } from "./components/Login";
import WebPlayback from "./components/WebPlayback";

function App() {
  const [token, setToken] = useState('');


  useEffect(() => {

    async function getToken() {
      const response = await fetch('http://localhost:8080/auth/token?user_id=0');
      const json = await response.json();
      console.log("response: " + json);
      if (json.access_token === undefined) {
        console.log("undefined");
        setToken("");
      } else {
        setToken(json.access_token);
      }
    }

    getToken();
    // console.log("test")

  }, []);


  return (
    <>
      {(token === "") ? <Login /> : <WebPlayback token={token} />}
    </>
  );
}

export default App;
