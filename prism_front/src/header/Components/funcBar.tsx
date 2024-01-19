import styled from "styled-components";
import { useEffect } from "react";
import { useSelector, useDispatch } from "react-redux";

import BeforeLogin from "./beforeLogin";
import AfterLogin from "./afterLogin";
import { AppDispatch, RootState } from "../../app/store";
import { fetchUser } from "../../app/slices/user/user";

const FuncBar: React.FC = () => {
  const dispatch = useDispatch<AppDispatch>();
  const userInfo = useSelector((state: RootState) => state.user);

  useEffect(() => {
    const userLoginCookie = document.cookie.includes("user_login");
    if (userLoginCookie) {
      // fetchUserInfo 액션 디스패치
      dispatch(fetchUser());
    }
  }, [dispatch]);

  return (
    <Box>
      {userInfo.user_id != "" ? (
        <AfterLogin userID={userInfo.user_id} />
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
