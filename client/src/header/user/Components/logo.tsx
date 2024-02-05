import styled from "styled-components";
import logo from "./../../../assets/logo.png";
import { useNavigate } from "react-router-dom";

const Logo = () => {
  const navigate = useNavigate();
  const handleLogoImgClick = () => {
    navigate("/home");
  };
  return <LogoImg onClick={handleLogoImgClick} src={logo}></LogoImg>;
};

export default Logo;

const LogoImg = styled.img`
  cursor: pointer;
`;
