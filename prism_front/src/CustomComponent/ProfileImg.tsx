import styled from "styled-components";

interface ComponentProps {
  id: string;
}

const ProfileImage: React.FC<ComponentProps> = ({ id }) => {
  return (
    <ImgBox
      src={`http://localhost:8080/assets/images/profiles/${id}.jpg`}></ImgBox>
  );
};

export default ProfileImage;

const ImgBox = styled.img`
  height: 100%;
`;
