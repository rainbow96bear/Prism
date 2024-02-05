import axios from "./../../../../configs/AxiosConfig";
import { useEffect, useState } from "react";
import { useDispatch } from "react-redux";
import { useNavigate } from "react-router-dom";
import styled from "styled-components";
import { AppDispatch, RootState } from "../../../../app/store";
import { login } from "../../../../app/slices/admin/admin";
import { useSelector } from "react-redux";

const Root = () => {
  const dispatch = useDispatch<AppDispatch>();
  const loginResult = useSelector(
    (state: RootState) => state.adminReducer.admin_info
  );
  const done = useSelector((state: RootState) => state.adminReducer.done);
  const navigate = useNavigate();
  const [password, setPassword] = useState(""); // 입력한 비밀번호를 상태로 관리

  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  };

  const handleSubmit = () => {
    dispatch(login(password));
  };
  useEffect(() => {
    if (done) {
      if (loginResult?.id != "") {
        navigate(`/admin/home/${loginResult?.id}`);
      }
    }
  }, [loginResult, navigate]);

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

const Box = styled.div``;

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
