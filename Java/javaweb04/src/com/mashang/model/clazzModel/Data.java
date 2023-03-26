package com.mashang.model.clazzModel;

public class Data {
    public Data(Long classId, String classNo, String createTime, String name, String remark, String type) {
        this.classId = classId;
        this.classNo = classNo;
        this.createTime = createTime;
        this.name = name;
        this.remark = remark;
        this.type = type;
    }

    // 班级主键ID
    private Long classId;

    // 班级编号
    private String classNo;

    // 创建时间
    private String createTime;

    // 班级名称
    private String name;

    // 备注
    private String remark;

    // 班级类型
    private String type;

    // Setter
    public void setClassId(Long classId) {
        this.classId = classId;
    }

    public void setClassNo(String classNo) {
        this.classNo = classNo;
    }

    public void setCreateTime(String createTime) {
        this.createTime = createTime;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void setRemark(String remark) {
        this.remark = remark;
    }

    public void setType(String type) {
        this.type = type;
    }

    // Getter
    public Long getClassId() {
        return classId;
    }

    public String getClassNo() {
        return classNo;
    }

    public String getCreateTime() {
        return createTime;
    }

    public String getName() {
        return name;
    }

    public String getRemark() {
        return remark;
    }

    public String getType() {
        return type;
    }
}
