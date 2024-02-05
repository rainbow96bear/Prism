import styled from "styled-components";

import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useSelector } from "react-redux";
import { AppDispatch, RootState } from "../../../../app/store";
import { useDispatch } from "react-redux";
import { logout } from "../../../../app/slices/admin/admin";

const Home = () => {
  const admin_info = useSelector(
    (state: RootState) => state.adminReducer.admin_info
  );
  const done = useSelector((state: RootState) => state.adminReducer.done);
  const navigate = useNavigate();
  const dispatch = useDispatch<AppDispatch>();
  const logoutFunc = async () => {
    dispatch(logout());
  };
  useEffect(() => {
    if (done) {
      if (admin_info.id == "") {
        navigate("/admin");
      } else {
        navigate(`/admin/home/${admin_info?.id}`);
      }
    }
  }, [admin_info]);
  return (
    <Box>
      <InfoBox>
        <div>Admin ID : {admin_info?.id}</div>
        <div>Admin Rank : {admin_info?.rank}</div>
        <button onClick={logoutFunc}> 로그아웃 </button>
      </InfoBox>
    </Box>
  );
};

export default Home;

const Box = styled.div``;

const InfoBox = styled.div`
  width: 50%;
  height: 50%;
  font-size: 2rem;
  font-weight: bold;
`;
