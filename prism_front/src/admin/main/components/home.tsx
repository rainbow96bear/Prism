import styled from "styled-components";

import { Admin } from "../../../GlobalType/Admin";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

interface HomeProps {
  admin_info: Admin | null;
}

const Home: React.FC<HomeProps> = ({ admin_info }) => {
  const navigate = useNavigate();
  useEffect(() => {
    console.log(admin_info?.id);
    navigate(`localhost:3000/admin/home/${admin_info?.id}`);
  }, []);
  return (
    <div>
      <div>{admin_info?.id}</div>
      <div>{admin_info?.rank}</div>
    </div>
  );
};

export default Home;
