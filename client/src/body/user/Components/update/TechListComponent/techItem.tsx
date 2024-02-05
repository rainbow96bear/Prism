import styled from "styled-components";
import { FC } from "react";
import { IoMdCloseCircleOutline } from "react-icons/io";
import { CiCircleMinus, CiCirclePlus } from "react-icons/ci";
import { TechData } from "../../../../../GlobalType/Tech";

interface TechItemProps extends TechData {
  onRemove: () => void;
  onDecrement: () => void;
  onIncrement: () => void;
}

const TechItem: FC<TechItemProps> = ({
  tech_name,
  level,
  onRemove,
  onDecrement,
  onIncrement,
}) => {
  const maxLevel = 10;

  return (
    <Container>
      <CloseIcon onClick={onRemove} />
      <Title>{tech_name}</Title>
      <MinusIcon onClick={onDecrement} />
      <LevelBar>
        {Array.from({ length: maxLevel }, (_, index) =>
          index < level ? (
            <LevelFillBlock key={index} />
          ) : (
            <LevelEmptyBlock key={index} />
          )
        )}
      </LevelBar>
      <PlusIcon onClick={onIncrement} />
    </Container>
  );
};

export default TechItem;

const Container = styled.div`
  display: flex;
  padding: 30px 0px;
  align-items: center;
`;

const Title = styled.div`
  background-color: hsl(0, 0%, 90%);
  border: 2px solid lightgray;
  border-radius: 20px;
  font-weight: bold;
  padding: 10px 15px;
  width: 110px;
  text-align: center;
`;

const LevelBar = styled.div`
  display: flex;
`;

const LevelEmptyBlock = styled.div`
  width: 20px;
  height: 20px;
  background-color: transparent;
  border: 2px solid hsl(999, 70%, 50%);
  margin-right: 5px;
`;

const LevelFillBlock = styled.div`
  width: 20px;
  height: 20px;
  background-color: hsl(999, 70%, 50%);
  border: 2px solid hsl(999, 70%, 50%);
  margin-right: 5px;
`;

const iconStyles = `
  font-size: 24px;
  margin: 0 10px;
  cursor: pointer;
`;

const CloseIcon = styled(IoMdCloseCircleOutline)`
  ${iconStyles}
  color: red;
`;

const MinusIcon = styled(CiCircleMinus)`
  ${iconStyles}
  color: black;
`;

const PlusIcon = styled(CiCirclePlus)`
  ${iconStyles}
  color: black;
`;
