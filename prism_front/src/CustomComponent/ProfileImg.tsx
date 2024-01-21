import styled from "styled-components";

interface ComponentProps {
  id: string;
}

const ProfileImage: React.FC<ComponentProps> = ({ id }) => {
  const imageUrl = `http://localhost:8080/assets/images/profiles/${id}.jpg`;

  return <ImgBox src={imageUrl}></ImgBox>;
};

export default ProfileImage;

const ImgBox = styled.img`
  height: 100%;
`;
