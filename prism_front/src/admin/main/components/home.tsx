import styled from "styled-components";

import { Admin } from "../../../GlobalType/Admin";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

interface HomeProps {
  admin_info: Admin | null;
}

const Home: React.FC<HomeProps> = ({ admin_info }) => {
  const navigate = useNavigate();
  const logout = async () => {
    try {
      await axios.get("http://localhost:8080/admin/user/logout", {
        withCredentials: true,
      });
      navigate("/admin");
    } catch (err) {
      console.log(err);
    }
  };
  useEffect(() => {
    navigate(`/admin/home/${admin_info?.id}`);
  }, []);
  return (
    <Box>
      <InfoBox>
        <div>Admin ID : {admin_info?.id}</div>
        <div>Admin Rank : {admin_info?.rank}</div>
        <button onClick={logout}> 로그아웃 </button>
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
