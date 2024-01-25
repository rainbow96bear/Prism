import styled from "styled-components";

interface ComponentProps {
  id: string;
}

const ProfileImage: React.FC<ComponentProps> = ({ id }) => {
  const server = process.env.REACT_APP_SERVER_URL;
  const imageUrl = `${server}/assets/images/profiles/${id}.jpg`;
  return <ImgBox src={imageUrl}></ImgBox>;
};

export default ProfileImage;

const ImgBox = styled.img`
  height: 100%;
`;
