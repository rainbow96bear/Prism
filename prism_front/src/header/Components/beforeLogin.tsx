import styled from "styled-components";

import loginImg from "./../../assets/kakao_login_small.png";

const BeforeLogin: React.FC = () => {
  // oauth 요청 URL
  const handleLogin = () => {
    window.location.href = "http://localhost:8080/kakao/login";
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
