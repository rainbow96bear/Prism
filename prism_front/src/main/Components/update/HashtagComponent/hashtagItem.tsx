import styled from "styled-components";

const HashTagItem = ({
  content,
  onRemove,
}: {
  content: string;
  onRemove: () => void;
}) => {
  return (
    <Container>
      <HashTagLogo>#</HashTagLogo>
      {content}
      <DeleteButton onClick={onRemove}>X</DeleteButton>
    </Container>
  );
};

export default HashTagItem;

const Container = styled.div`
  display: flex;
  align-items: center;
  font-size: 1.2rem;
`;
const HashTagLogo = styled.div`
  padding: 5px;
  font-weight: bold;
`;
const DeleteButton = styled.div`
  padding: 5px;
  color: #aaaaaa;
  cursor: pointer;
`;
