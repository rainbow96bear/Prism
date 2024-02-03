import styled from "styled-components";
import { useEffect } from "react";
import { useSelector, useDispatch } from "react-redux";

import BeforeLogin from "./beforeLogin";
import AfterLogin from "./afterLogin";
import { AppDispatch, RootState } from "../../app/store";
import { getUserInfo } from "../../app/slices/user/user";

const FuncBar: React.FC = () => {
  const dispatch = useDispatch<AppDispatch>();
  const userInfo = useSelector((state: RootState) => state.user);

  useEffect(() => {
    const userLoginCookie = document.cookie.includes("User");
    if (userLoginCookie) {
      // fetchUserInfo 액션 디스패치 여기에 id 값을 넣어야해 여기는 userInfo
      dispatch(getUserInfo());
    }
  }, [dispatch, userInfo]);
  return (
    <Box>
      {userInfo.id != "" ? (
        <AfterLogin userID={userInfo.id} />
      ) : (
        <BeforeLogin />
      )}
    </Box>
  );
};

export default FuncBar;

const Box = styled.div`
  display: flex;
  justify-content: right;
  height: 100%;
  align-items: center;
`;
