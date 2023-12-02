import styled from "styled-components";

import loginImg from "./../../assets/kakao_login_small.png";

const AfterLogin: React.FC = () => {
  // oauth 요청 URL
  const handleLogin = () => {
    window.location.href = "http://localhost:8080/kakaoLogin";
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
