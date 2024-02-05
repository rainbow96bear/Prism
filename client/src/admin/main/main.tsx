import styled from "styled-components";
import { Route, Routes, useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";
import axios from "./../../configs/AxiosConfig";

import { Admin } from "../../GlobalType/Admin";
import Root from "./components/root";
import Home from "./components/home";
import Loading from "../../CustomComponent/Loading";
import AdminHeader from "../header/header";
import Setting from "./settingComponent/setting";

const AdminMain = () => {
  const [admin_info, setAdmin_info] = useState<Admin | null>(null);
  const [isAdmin, setIsAdmin] = useState(false);

  const navigate = useNavigate();
  useEffect(() => {
    const originURL = window.location.pathname;
    const checkAdmin = async () => {
      try {
        const checkResult = (
          await axios.get("/admins/authorization", {
            withCredentials: true,
          })
        ).data;
        if (checkResult?.isAdmin == false) {
          navigate("/");
        } else if (checkResult?.admin_info.id == "") {
          navigate("/admin");
        } else if (checkResult?.admin_info.id != "") {
          if (originURL == "/admin") {
            navigate("/admin/home");
          } else {
            navigate(originURL);
          }
        }
        setAdmin_info(checkResult?.admin_info);
        setIsAdmin(checkResult?.isAdmin);
      } catch (error) {
        console.error("Admin 확인 중 에러 발생:", error);
      }
    };

    checkAdmin();
  }, [window.location.pathname]);

  return (
    <>
      {admin_info?.id != null ? (
        <HeaderArea>
          <AdminHeader></AdminHeader>
        </HeaderArea>
      ) : (
        <></>
      )}
      <BodyArea>
        <Box>
          <Routes>
            <Route path="/" element={<Root setAdmin_info={setAdmin_info} />} />
            <Route path="/home/*" element={<Home admin_info={admin_info} />} />
            <Route path="/setting/*" element={<Setting></Setting>} />
          </Routes>
        </Box>
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
