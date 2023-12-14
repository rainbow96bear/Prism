import axios from "axios";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import styled from "styled-components";

interface setAdmin_info {
  setAdmin_info: Function;
}

const Root: React.FC<setAdmin_info> = ({ setAdmin_info }) => {
  const navigate = useNavigate();
  const [password, setPassword] = useState(""); // 입력한 비밀번호를 상태로 관리

  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  };

  const handleSubmit = async () => {
    try {
      const loginResult = await axios.post(
        "http://localhost:8080/admin/user/login",
        { password: password } // 비밀번호를 객체 형태로 body에 담아 전송
      );
      setAdmin_info(loginResult.data);
      navigate(`http://localhost:3000/admin/home`);
    } catch (error) {
      console.error("Admin 정보를 가져오는 중 에러 발생:", error);
    }
  };

  useEffect(() => {
    const checkAdmin = async () => {
      try {
        const loginResult = await axios.post(
          "http://localhost:8080/admin/user/login"
        );
        setAdmin_info(loginResult.data);
        navigate(`http://localhost:3000/admin/home`);
      } catch (error) {
        console.error("Admin 정보를 가져오는 중 에러 발생:", error);
      }
    };

    checkAdmin();
  }, []);

  return (
    <Box>
      <DescribeBox>관리자 접속</DescribeBox>
      <PasswordBox>
        <input
          type="password"
          value={password}
          onChange={handlePasswordChange}
        />
        <button onClick={handleSubmit}>접속</button>
      </PasswordBox>
    </Box>
  );
};

export default Root;

const Box = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 50vh;
`;

const DescribeBox = styled.div`
  font-size: 20px;
  font-weight: bold;
`;

const PasswordBox = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  input {
    padding: 0px;
    margin: 2px;
    font-size: 25px;
  }
  button {
    padding: 0px;
    margin: 2px;
    font-size: 20px;
    font-weight: bold;
  }
`;
