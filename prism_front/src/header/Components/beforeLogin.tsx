import styled from "styled-components";

import loginImg from "./../../assets/kakao_login_small.png";

const BeforeLogin: React.FC = () => {
  // oauth 요청 URL
  const handleLogin = async () => {
    const client_id = process.env.REACT_APP_REST_API_KEY;
    const redirect_uri = process.env.REACT_APP_REDIRECT_URI;
    window.location.href = `https://kauth.kakao.com/oauth/authorize?response_type=code&client_id=${client_id}&redirect_uri=${redirect_uri}`;
  };

  https: return (
    <>
      <ButtomBox onClick={handleLogin}>
        <LoginImg src={loginImg} />
      </ButtomBox>
    </>
  );
};

export default BeforeLogin;

const ButtomBox = styled.div`
  height: 50%;
  margin: 0px 10px;
  cursor: pointer;
`;

const LoginImg = styled.img`
  height: 100%;
`;
