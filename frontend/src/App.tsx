import { useEffect, useState } from "react";

import "./App.css";
// import Home from "./components/Home";
import { Login } from "./components/Login";
import WebPlayback from "./components/WebPlayback";

function App() {
  const [token, setToken] = useState('');

  // call to get token on app startup

  useEffect(() => {

    async function getToken() {
      const response = await fetch('http://localhost:8080/auth/token');
      const json = await response.json();
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
