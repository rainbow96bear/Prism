import styled from "styled-components";

interface AfterLogin {
  type: string;
}

const darkSkyBlue = "rgba(0, 102, 204, 0.8)";

const AfterLogin: React.FC = () => {
  const Rest_api_key = "b56eb8d850568962a66cb7b117530c3c"; //REST API KEY
  const redirect_uri = "http://localhost:3000/auth"; //Redirect URI
  // oauth 요청 URL
  const kakaoURL = `https://kauth.kakao.com/oauth/authorize?client_id=${Rest_api_key}&redirect_uri=${redirect_uri}&response_type=code`;
  const handleLogin = () => {
    window.location.href = kakaoURL;
  };
  https: return (
    <>
      <button onClick={handleLogin}>카카오 로그인</button>
      <ButtomBox type="login">로그인</ButtomBox>
      <ButtomBox type="signUp">회원가입</ButtomBox>
    </>
  );
};

export default AfterLogin;

const ButtomBox = styled.div<AfterLogin>`
  height: 50%;
  margin: 0px 10px;
  cursor: pointer;
  font-weight: bold;
  padding: 3px 7px;
  border-radius: 7px;
  background-color: ${({ type }) =>
    type === "login" ? "lightblue" : darkSkyBlue};
`;
