import { useEffect, useState } from "react";

import "./App.css";
// import Home from "./components/Home";
import { Login } from "./components/Login";
import WebPlayback from "./components/WebPlayback";

function App() {
  const [token, setToken] = useState('');


  useEffect(() => {

    async function getToken() {
      const response = await fetch('http://localhost:8080/auth/token');
      const json = await response.json();
      console.log(json);
      setToken(json.access_token);
    }

    getToken();

  }, []);


  return (
    <>
      {(token === "") ? <Login /> : <WebPlayback token={token} />}
    </>
  );
}

export default App;
