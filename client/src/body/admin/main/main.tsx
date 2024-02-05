import styled from "styled-components";
import { Route, Routes, useNavigate } from "react-router-dom";
import { useEffect } from "react";
import { useSelector, useDispatch } from "react-redux";
import { getAdminAuth } from "../../../app/slices/admin/admin";
import { AppDispatch, RootState } from "../../../app/store";
import Root from "./components/root";
import Home from "./components/home";
import Setting from "./settingComponent/setting";
import Loading from "../../../CustomComponent/Loading";

const AdminMain = () => {
  const dispatch = useDispatch<AppDispatch>();
  const navigate = useNavigate();
  const isAdmin = useSelector((state: RootState) => state.adminReducer.isAdmin);
  const done = useSelector((state: RootState) => state.adminReducer.done);
  const admin_info = useSelector(
    (state: RootState) => state.adminReducer.admin_info
  );
  useEffect(() => {
    try {
      dispatch(getAdminAuth());
    } catch (error) {
      console.error("Admin 정보를 가져오는 중 에러 발생:", error);
    }
  }, [dispatch]);
  useEffect(() => {
    if (done) {
      const originURL = window.location.pathname;
      if (isAdmin === false) {
        navigate("/");
      } else if (admin_info.id === "") {
        navigate("/admin");
      } else if (admin_info.id !== "") {
        if (originURL === "/admin") {
          navigate("/admin/home");
        } else {
          navigate(originURL);
        }
      }
    }
  }, [isAdmin, admin_info, navigate]);
  return (
    <>
      <BodyArea>
        {done ? (
          <Box>
            <Routes>
              <Route path="/" element={<Root />} />
              <Route path="/home/*" element={<Home />} />
              <Route path="/setting/*" element={<Setting />} />
            </Routes>
          </Box>
        ) : (
          <Loading></Loading>
        )}
      </BodyArea>
    </>
  );
};

export default AdminMain;

const HeaderArea = styled.div`
  display: flex;
  justify-content: center;
  width: 100%;
  height: 70px;
  border-bottom: 2px solid gray;
`;

const BodyArea = styled.div`
  width: 100%;
  padding-top: 50px;
  min-height: calc(100vh - 70px);
  border: 1px solid lightgray;
  border-top: none;
  border-bottom: none;
`;

const Box = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  > div {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    width: 70%;
    height: 100%;
  }
`;
