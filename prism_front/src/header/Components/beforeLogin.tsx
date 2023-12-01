import styled from "styled-components";

import loginImg from "./../../assets/kakao_login_small.png";

const AfterLogin: React.FC = () => {
  const Rest_api_key = process.env.REACT_APP_REST_API_KEY;
  const redirect_uri = process.env.REACT_APP_REDIRECT_URI;
  // oauth 요청 URL
  const kakaoURL = `https://kauth.kakao.com/oauth/authorize?client_id=${Rest_api_key}&redirect_uri=${redirect_uri}&response_type=code`;
  const handleLogin = () => {
    window.location.href = kakaoURL;
    console.log(window.location.href);
  };
  https: return (
    <>
      <ButtomBox onClick={handleLogin}>
        <img src={loginImg} />
      </ButtomBox>
    </>
  );
};

export default AfterLogin;

const ButtomBox = styled.div`
  height: 50%;
  margin: 0px 10px;
  cursor: pointer;
`;
