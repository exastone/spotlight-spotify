// AuthenticationComponent.tsx
import React, { useEffect, useState } from 'react';

interface AuthenticationProps {
  childPlayer: React.ReactNode;
}

const Authentication: React.FC<AuthenticationProps> = ({ childPlayer: children }) => {
  const [accessToken, setAccessToken] = useState<string | null>(null);

  useEffect(() => {
    // Fetch the access token from the backend here
    // Store the token securely using React context or state management (e.g., Redux)
    // Update the accessToken state with the retrieved token
    async function getToken() {
      const response = await fetch('http://localhost:8080/auth/token');
      const json = await response.json();
      setAccessToken(json.access_token);
    }

  }, []);

  return (
    <>
      {accessToken ? (
        // Render children components once the access token is available
        children
      ) : (
        // Render a loading or authentication UI while waiting for the token
        <div>Loading...</div>
      )}
    </>
  );
};

export default Authentication;

/* 
Authorization Code Flow

The authorization code flow is suitable for long-running applications (e.g. web and mobile apps) where the user grants permission only once.

If you’re using the authorization code flow in a mobile app, or any other type of application where the client secret can't be safely stored, then you should use the PKCE extension. Keep reading to learn how to correctly implement it.

The following diagram shows how the authorization code flow works:

[REQUEST-1] : [Application] -> [Spotify Accounts Service]
  Description: Request authorization to access user data
  
  GET Request:
    Endpoint: /authorize
    QUERY parameters:
      client_id
      response_type="code"
      redirect_uri
      state (optional)
      scope

  [Spotify Accounts Service] -> [User] 
    Description: User is prompted to login and authorize access to data by application

    If user authorizes access, then:
      User is redirected to *redirect_uri* specified in App setting (Spotify Account Dashboard),
      returning user back to the application, triggering response.
  
  [RESPONSE-1] : [Application] <- [User]
    Description: Response sent from Spotify Accounts Service to Application
    
    QUERY parameters:
      code - authorization code (to be exchnaged for access token)
      state - value of the state parameter supplied in the request.

[REQUEST-2] : [Application] -> [Spotify Accounts Service]
  Description: Request Access Token using auth. code
  
  POST Request:
    Endpoint: /api/token
    BODY Parameters: (application/x-www-form-urlencoded)
      grant_type="authorization_code"
      code - authorization code returned from the previous request
      redirect_uri - used for validation only (no redirection occurs).
        Value must match the value of redirect_uri supplied to auth. code request (req. 1).
    
    HEADER Parameters:
      Authorization - Basic <base64 encoded client_id:client_secret>
      Content-Type - application/x-www-form-urlencoded

  [RESPONSE-2] : [Application] <- [Spotify Accounts Service]
    Description: response body contains access and refresh token as JSON data
    JSON data:
      access_token - access token for API access
      token_type - "Bearer"
      scope - list of scopes granted by user associated with access token
      expires_in - 3600 (seconds)
      refresh_token - token that can be used to request a new access token.

[REQUEST-3] : [Application] -> [Spotify Accounts Service]
  Description: request a refresh access_token:
  
  POST Request:
    Endpoint: /api/token
    BODY Parameters: (application/x-www-form-urlencoded)
      grant_type="refresh_token"
      refresh_token - refresh token returned from inital authorization code exchange
    
    HEADER Parameters:
      Authorization - Basic <base64 encoded client_id:client_secret>
      Content-Type - application/x-www-form-urlencoded
  
  [RESPONSE-3] : [Application] <- [Spotify Accounts Service]
    Description: response body contains new access_token as JSON data
    JSON data:
      access_token - access token for API access
      token_type - "Bearer"
      scope - list of scopes granted by user associated with access token
      expires_in - 3600 (seconds)
  
      refresh_token ? Docs say "A new refresh token might be returned too"?


*/