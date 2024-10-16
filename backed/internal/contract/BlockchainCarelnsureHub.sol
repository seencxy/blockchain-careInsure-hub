// SPDX-License-Identifier: MIT
pragma solidity >=0.6.0;

pragma experimental ABIEncoderV2;

contract RetirementInsurance {
    // 定义保险套餐
    struct Package {
        uint256 id;
        uint256 price;
        uint256 monthlyPension;
        uint256 start;
        uint256 coverageAge;
        bool highMedicalCoverage;
        uint256 refundPeriod;
        bool isActive; // 新增字段，表示套餐是否可用
    }

    // 定义用户
    struct User {
        string username; // 用户姓名
        uint256 year; // 生日日期
        uint256 month;
        uint256 buyDate; // 购买日期
        uint256 kind; // 套餐类型
    }
    // 保险套餐映射
    mapping(uint256 => Package) public packages;
    // 存放所有购买人员
    User[] public buyPackageUsers;

    // 保险套餐购买数量映射
    mapping(uint256 => uint256) public packagePurchaseCounts;


    // 合约管理员
    address public owner;
    uint256 public nextPackageId = 1;

    constructor() {
        owner = msg.sender;
    }

    // 允许所有者添加新的保险套餐
    function addPackage(
        uint256 price,
        uint256 monthlyPension,
        uint256 start,
        uint256 coverageAge,
        bool highMedicalCoverage,
        uint256 refundPeriod
    ) public onlyOwner {
        packages[nextPackageId] = Package(
            nextPackageId,
            price,
            monthlyPension,
            start,
            coverageAge,
            highMedicalCoverage,
            refundPeriod,
            true
        );
        nextPackageId++;
    }

    // 确保只有管理员能调用
    modifier onlyOwner() {
        require(msg.sender == owner, "Not authorized");
        _;
    }

    // 返回所有保险套餐的详细信息
    function getAllPackages() public view returns (Package[] memory) {
        Package[] memory allPackages = new Package[](nextPackageId - 1);
        for (uint256 i = 1; i < nextPackageId; i++) {
            allPackages[i - 1] = packages[i];
        }
        return allPackages;
    }

    // 检查并返回用户未购买的所有套餐ID列表
    function getUnpurchasedPackages(string memory username) public view returns (uint256[] memory) {
        uint256[] memory unpurchased = new uint256[](nextPackageId - 1);
        uint256 count = 0;

        // 初始化一个数组来标记所有套餐的购买状态，初始值为false
        bool[] memory purchased = new bool[](nextPackageId);

        // 标记用户已购买的套餐
        for (uint256 i = 0; i < buyPackageUsers.length; i++) {
            if (keccak256(abi.encodePacked(buyPackageUsers[i].username)) == keccak256(abi.encodePacked(username))) {
                purchased[buyPackageUsers[i].kind] = true;
            }
        }

        // 遍历所有套餐，找出未购买的套餐
        for (uint256 i = 1; i < nextPackageId; i++) {
            if (!purchased[i]) {
                unpurchased[count] = i;
                count++;
            }
        }

        // 调整数组大小以匹配未购买套餐的实际数量
        uint256[] memory result = new uint256[](count);
        for (uint256 i = 0; i < count; i++) {
            result[i] = unpurchased[i];
        }

        return result;
    }


    // 购买保险套餐
    function purchasePackage(
        string memory username,
        uint256 year,
        uint256 month,
        uint256 packageId
    ) public {
        require(packages[packageId].price > 0 && packages[packageId].isActive, "Package not defined or not active.");

        // 检查用户是否已经购买过同类型的套餐
        for (uint256 i = 0; i < buyPackageUsers.length; i++) {
            if (
                keccak256(abi.encodePacked(buyPackageUsers[i].username)) ==
                keccak256(abi.encodePacked(username)) &&
                buyPackageUsers[i].kind == packageId
            ) {
                require(
                    buyPackageUsers[i].kind != packageId,
                    "User has already purchased this package type"
                );
                break;
            }
        }

        // 创建用户实例
        User memory newUser = User(
            username,
            year,
            month,
            block.timestamp,
            packageId
        );
        // 将用户添加到购买人员名单
        buyPackageUsers.push(newUser);
        // 增加套餐的购买数量
        packagePurchaseCounts[packageId]++;
    }

    // 获取购买人员信息
    function getBuyerInfo(string memory username)
    public
    view
    returns (User[] memory)
    {
        uint256 count = 0;

        // 首先计算匹配的用户数量
        for (uint256 i = 0; i < buyPackageUsers.length; i++) {
            if (
                keccak256(abi.encodePacked(buyPackageUsers[i].username)) ==
                keccak256(abi.encodePacked(username))
            ) {
                count++;
            }
        }

        // 根据计数创建新数组
        User[] memory matchingUsers = new User[](count);
        uint256 index = 0;

        // 再次遍历以填充数组
        for (uint256 i = 0; i < buyPackageUsers.length; i++) {
            if (
                keccak256(abi.encodePacked(buyPackageUsers[i].username)) ==
                keccak256(abi.encodePacked(username))
            ) {
                matchingUsers[index] = buyPackageUsers[i];
                index++;
            }
        }

        return matchingUsers;
    }

    // 取消保险套餐
    function cancelPackage(string memory username, uint256 packageId) public {
        // 查找用户及其购买的套餐
        for (uint256 i = 0; i < buyPackageUsers.length; i++) {
            if (
                keccak256(abi.encodePacked(buyPackageUsers[i].username)) ==
                keccak256(abi.encodePacked(username)) &&
                packageId == buyPackageUsers[i].kind
            ) {
                User storage user = buyPackageUsers[i];
                Package storage pkg = packages[user.kind];

                uint256 refundPeriodSeconds = pkg.refundPeriod * 30 * 86400;
                // 检查是否在退款期限内
                require(
                    block.timestamp <= user.buyDate + refundPeriodSeconds,
                    "Refund period expired"
                );

                // 处理退款逻辑（如果有的话）
                // 例如，将保险费用退回给用户（这里需要根据实际情况进行实现）

                // 移除用户购买信息
                delete buyPackageUsers[i];
                return;
            }
        }
        revert("User not found or already cancelled");
    }

    // 定义一个新结构体用于返回用户信息和套餐ID
    struct UserInfo {
        string username;
        uint256 packageId;
    }

    // 新函数：获取所有购买人员的用户名和套餐ID
    function getAllBuyersInfo() public view returns (UserInfo[] memory) {
        UserInfo[] memory allUsersInfo = new UserInfo[](buyPackageUsers.length);

        for (uint256 i = 0; i < buyPackageUsers.length; i++) {
            // 将用户名和套餐ID添加到新数组
            allUsersInfo[i] = UserInfo({
                username: buyPackageUsers[i].username,
                packageId: buyPackageUsers[i].kind
            });
        }

        return allUsersInfo;
    }

    // 下面实现一个方法 检查用户年龄是否到了 保险年龄 如果到了则触发一个需要发养老金的事件
    // 定义一个事件，用于通知养老金发放
    event PensionPaid(string username, uint256 amount);

    // 由后端服务调用的发放养老金函数
    function payPensions(uint256 currentYear, uint256 currentMonth)
    public
    onlyOwner
    {
        for (uint256 i = 0; i < buyPackageUsers.length; i++) {
            User storage user = buyPackageUsers[i];
            Package storage pkg = packages[user.kind];
            uint256 old = currentYear - user.year;
            if (currentMonth <= user.month) {
                old = old + 1;
            }
            if (old >= pkg.start && old <= pkg.coverageAge) {
                // 发放养老金逻辑
                uint256 pensionAmount = pkg.monthlyPension;
                // 触发事件
                emit PensionPaid(user.username, pensionAmount);
            }
        }
    }

    function getPackagePurchaseCount(uint256 packageId) public view returns (uint256) {
        return packagePurchaseCounts[packageId];
    }

    function updatePackageStatus(uint256 packageId) public onlyOwner {
        packages[packageId].isActive = !packages[packageId].isActive;
    }

}
