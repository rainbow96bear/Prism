import React, { useEffect, useState } from "react";
import styled from "styled-components";
import Table from "./table";
import axios from "./../../../../configs/AxiosConfig";
import AddTechModal from "./modal";

// Tech 컴포넌트

const Tech = () => {
  const column = ["코드", "기술 명", "Count"];
  const [selectedOption, setSelectedOption] = useState("전체");
  const [info, setInfo] = useState([[]]);
  const [sortedInfo, setSortedInfo] = useState([[]]);
  const [searchKeyword, setSearchKeyword] = useState(""); // 추가
  const [isModalOpen, setModalOpen] = useState(false);

  const openModal = () => {
    setModalOpen(true);
  };

  const closeModal = () => {
    setModalOpen(false);
  };
  const getTechList = async () => {
    const list = (
      await axios.get("/admin/access/tech", {
        withCredentials: true,
      })
    ).data;
    setInfo(list);
    setSortedInfo(list);
  };

  useEffect(() => {
    getTechList();
  }, []);

  const handleSelectChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedValue = e.target.value;
    setSelectedOption(selectedValue);

    if (selectedValue === "전체") {
      setSortedInfo(info);
    } else {
      const columnIndex = column.indexOf(selectedValue);
      setSortedInfo((prevSortedInfo) =>
        [...prevSortedInfo].sort((a, b) => {
          const aValue = a[columnIndex] as string | undefined;
          const bValue = b[columnIndex] as string | undefined;

          if (aValue !== undefined && bValue !== undefined) {
            return aValue.localeCompare(bValue);
          } else {
            return 0;
          }
        })
      );
    }
  };

  const handleSearch = () => {
    // 검색어에 해당하는 결과만 필터링
    const filteredInfo = info.filter((row) => {
      const rowValues = Object.values(row);
      return rowValues.some((cell, cellIndex) => {
        if (
          selectedOption === "Count" &&
          cellIndex === column.indexOf("Count")
        ) {
          // "Count" 열에 대한 검색일 때는 문자열로 변환하지 않고 그대로 비교
          return String(cell) === searchKeyword;
        } else {
          // 다른 열에 대한 검색은 문자열로 변환하여 비교
          return String(cell)
            .toLowerCase()
            .includes(searchKeyword.toLowerCase());
        }
      });
    });
    setSortedInfo(filteredInfo);
  };

  return (
    <Box>
      <FuncBar>
        <SearchBar>
          <select value={selectedOption} onChange={handleSelectChange}>
            <option value="전체">전체</option>
            {column.map((data, idx) => (
              <option key={"column_" + idx} value={data}>
                {data}
              </option>
            ))}
          </select>
          <input
            value={searchKeyword}
            onChange={(e) => setSearchKeyword(e.target.value)}
          />
          <button onClick={handleSearch}>검색</button>
        </SearchBar>
        <button onClick={openModal}>추가</button>
      </FuncBar>
      <Table column={column} info={sortedInfo} getTechList={getTechList} />
      <AddTechModal
        isOpen={isModalOpen}
        onRequestClose={closeModal}
        getTechList={getTechList}></AddTechModal>
    </Box>
  );
};

export default Tech;

const Box = styled.div`
  width: 100%;
`;

const FuncBar = styled.div`
  display: flex;
  justify-content: space-between;
  padding-bottom: 30px;
  height: 50px;
  & button {
    padding: 0px 10px;
  }
`;

const SearchBar = styled.div`
  height: 50px;
  display: flex;
  align-items: center;

  & select,
  & input,
  & button {
    height: 100%;
    padding: 0;
    padding: 0px 10px;
    margin-right: 10px;
  }
`;
