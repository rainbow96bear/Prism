import styled from "styled-components";
import { Route, Routes } from "react-router-dom";
import { useEffect, useState } from "react";
import axios from "axios";

import { Admin } from "../../GlobalType/Admin";
import Root from "./components/root";
import Home from "./components/home";

const AdminMain = () => {
  const [admin_info, setAdmin_info] = useState<Admin | null>(null);
  useEffect(() => {
    const checkAdmin = async () => {
      try {
        const checkResult = await axios.get(
          "http://localhost:8080/admin/user/check"
        );
        if (checkResult.data.isAdmin == "false") {
          window.location.href = "http://localhost:3000/";
        }
      } catch (error) {
        console.error("Admin 확인 중 에러 발생:", error);
      }
    };

    checkAdmin();
  }, []);

  return (
    <div>
      <Routes>
        <Route path="/" element={<Root setAdmin_info={setAdmin_info} />} />
        <Route path="/home" element={<Home admin_info={admin_info} />} />
      </Routes>
    </div>
  );
};

export default AdminMain;
