import styled from "styled-components";
import { useLocation, useNavigate } from "react-router-dom";

interface CommunityProps {
  isCommunity: boolean;
}

interface ProjectProps {
  isProject: boolean;
}

const Category: React.FC = () => {
  const location = useLocation();
  const navigate = useNavigate();

  const isCommunity: boolean = location.pathname === "/community";
  const isProject: boolean = location.pathname === "/project";

  const handleCommunityClick = () => {
    navigate("/community");
  };
  const handleProjectClick = () => {
    navigate("/project");
  };
  return (
    <Box>
      <Community onClick={handleCommunityClick} isCommunity={isCommunity}>
        커뮤니티
      </Community>
      <Project onClick={handleProjectClick} isProject={isProject}>
        프로젝트
      </Project>
    </Box>
  );
};

const Box = styled.div`
  display: flex;
  font-size: 1.5rem;
  height: 100%;
  div {
    margin: 0px 10px 0px 10px;
    display: flex;
    align-items: center;
    cursor: pointer;
  }
`;

const Community = styled.div<CommunityProps>`
  color: ${({ isCommunity }) => (isCommunity ? "black" : "lightgray")};
  border-bottom: ${({ isCommunity }) => (isCommunity ? "4px solid green" : "")};
`;
const Project = styled.div<ProjectProps>`
  color: ${({ isProject }) => (isProject ? "black" : "lightgray")};
  border-bottom: ${({ isProject }) => (isProject ? "4px solid green" : "")};
`;

export default Category;
