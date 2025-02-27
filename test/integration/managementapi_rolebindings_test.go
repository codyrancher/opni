package integration_test

import (
	"context"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	corev1 "github.com/rancher/opni/pkg/apis/core/v1"
	managementv1 "github.com/rancher/opni/pkg/apis/management/v1"
	"github.com/rancher/opni/pkg/test"
)

// #region Test Setup
var _ = Describe("Management API Rolebinding Management Tests", Ordered, Label("integration"), func() {
	var environment *test.Environment
	var client managementv1.ManagementClient
	BeforeAll(func() {
		environment = &test.Environment{}
		Expect(environment.Start()).To(Succeed())
		client = environment.NewManagementClient()

		_, err := client.CreateRole(context.Background(), &corev1.Role{
			Id: "test-role",
		},
		)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterAll(func() {
		Expect(environment.Stop()).To(Succeed())
	})
	//#endregion

	//#region Happy Path Tests

	var err error
	When("creating a new rolebinding", func() {

		It("can get information about all rolebindings", func() {
			_, err = client.CreateRoleBinding(context.Background(), &corev1.RoleBinding{
				Id:       "test-rolebinding1",
				RoleId:   "test-role",
				Subjects: []string{"test-subject"},
			})
			Expect(err).NotTo(HaveOccurred())

			rolebindingInfo, err := client.GetRoleBinding(context.Background(), &corev1.Reference{
				Id: "test-rolebinding1",
			})
			Expect(err).NotTo(HaveOccurred())

			Expect(rolebindingInfo.Id).To(Equal("test-rolebinding1"))
			Expect(rolebindingInfo.RoleId).To(Equal("test-role"))
			Expect(rolebindingInfo.Subjects).To(Equal([]string{"test-subject"}))
		})
	})

	It("can list all rolebindings", func() {
		rolebinding, err := client.ListRoleBindings(context.Background(), &emptypb.Empty{})
		Expect(err).NotTo(HaveOccurred())

		rolebindingList := rolebinding.Items
		Expect(rolebindingList).To(HaveLen(1))
		for _, rolebindingItem := range rolebindingList {
			Expect(rolebindingItem.Id).To(Equal("test-rolebinding1"))
			Expect(rolebindingItem.RoleId).To(Equal("test-role"))
			Expect(rolebindingItem.Subjects).To(ContainElement("test-subject"))
		}
	})

	It("can delete an existing rolebinding", func() {
		_, err = client.DeleteRoleBinding(context.Background(), &corev1.Reference{
			Id: "test-rolebinding1",
		})
		Expect(err).NotTo(HaveOccurred())

		_, err = client.GetRoleBinding(context.Background(), &corev1.Reference{
			Id: "test-rolebinding1",
		})
		Expect(err).To(HaveOccurred())
		Expect(status.Code(err)).To(Equal(codes.NotFound))
	})

	//#endregion

	//#region Edge Case Tests

	It("cannot create rolebindings without a RoleID", func() {
		_, err = client.CreateRoleBinding(context.Background(), &corev1.RoleBinding{
			Id:       "test-rolebinding2",
			Subjects: []string{"test-subject"},
		})
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("missing required field: roleId"))

		_, err = client.GetRoleBinding(context.Background(), &corev1.Reference{
			Id: "test-rolebinding2",
		})
		Expect(err).To(HaveOccurred())
		Expect(status.Code(err)).To(Equal(codes.NotFound))
	})

	It("can create rolebindings without a valid RoleId", func() {
		_, err = client.CreateRoleBinding(context.Background(), &corev1.RoleBinding{
			RoleId:   uuid.NewString(),
			Id:       "test-rolebinding2",
			Subjects: []string{"test-subject"},
		})
		Expect(err).NotTo(HaveOccurred())

		rbInfo, err := client.GetRoleBinding(context.Background(), &corev1.Reference{
			Id: "test-rolebinding2",
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(rbInfo.Taints).To(ContainElement("role not found"))
	})

	It("cannot create rolebindings without an Id", func() {
		_, err = client.CreateRoleBinding(context.Background(), &corev1.RoleBinding{
			RoleId:   "test-role",
			Subjects: []string{"test-subject"},
		})
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("missing required field: id"))
	})

	It("can create and get rolebindings without a subject", func() {
		_, err = client.CreateRoleBinding(context.Background(), &corev1.RoleBinding{
			Id:     "test-rolebinding3",
			RoleId: "test-role",
		})
		Expect(err).NotTo(HaveOccurred())

		rolebindingInfo, err := client.GetRoleBinding(context.Background(), &corev1.Reference{
			Id: "test-rolebinding3",
		})
		Expect(err).NotTo(HaveOccurred())

		Expect(rolebindingInfo.Id).To(Equal("test-rolebinding3"))
		Expect(rolebindingInfo.RoleId).To(Equal("test-role"))
		Expect(rolebindingInfo.Subjects).To(BeNil())
		Expect(rolebindingInfo.Taints).To(Equal([]string{"no subjects"}))

		_, err = client.DeleteRoleBinding(context.Background(), &corev1.Reference{
			Id: "test-rolebinding3",
		})
		Expect(err).NotTo(HaveOccurred())
	})

	It("can create and get rolebindings without a taint", func() {
		_, err = client.CreateRoleBinding(context.Background(), &corev1.RoleBinding{
			Id:       "test-rolebinding4",
			RoleId:   "test-role",
			Subjects: []string{"test-subject"},
		})
		Expect(err).NotTo(HaveOccurred())

		rolebindingInfo, err := client.GetRoleBinding(context.Background(), &corev1.Reference{
			Id: "test-rolebinding4",
		})
		Expect(err).NotTo(HaveOccurred())

		Expect(rolebindingInfo.Id).To(Equal("test-rolebinding4"))
		Expect(rolebindingInfo.RoleId).To(Equal("test-role"))
		Expect(rolebindingInfo.Subjects).To(ContainElement("test-subject"))
		Expect(rolebindingInfo.Taints).To(BeNil())

		_, err = client.DeleteRoleBinding(context.Background(), &corev1.Reference{
			Id: "test-rolebinding4",
		})
		Expect(err).NotTo(HaveOccurred())
	})

	It("cannot delete a rolebinding without specifying an Id", func() {
		_, err = client.CreateRoleBinding(context.Background(), &corev1.RoleBinding{
			Id:       "test-rolebinding5",
			RoleId:   "test-role",
			Subjects: []string{"test-subject"},
		})
		Expect(err).NotTo(HaveOccurred())

		_, err = client.DeleteRoleBinding(context.Background(), &corev1.Reference{})
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("missing required field: id"))

		_, err = client.DeleteRoleBinding(context.Background(), &corev1.Reference{
			Id: "test-rolebinding5",
		})
		Expect(err).NotTo(HaveOccurred())
	})

	It("cannot delete a rolebinding without specifying a valid Id", func() {
		_, err = client.CreateRoleBinding(context.Background(), &corev1.RoleBinding{
			Id:       "test-rolebinding6",
			RoleId:   "test-role",
			Subjects: []string{"test-subject"},
		})
		Expect(err).NotTo(HaveOccurred())

		_, err = client.DeleteRoleBinding(context.Background(), &corev1.Reference{
			Id: uuid.NewString(),
		})
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("not found"))

		_, err = client.DeleteRoleBinding(context.Background(), &corev1.Reference{
			Id: "test-rolebinding6",
		})
		Expect(err).NotTo(HaveOccurred())
	})

	It("cannot create rolebindings with identical Ids", func() {
		_, err = client.CreateRoleBinding(context.Background(), &corev1.RoleBinding{
			Id:       "test-rolebinding7",
			RoleId:   "test-role",
			Subjects: []string{"test-subject"},
		})
		Expect(err).NotTo(HaveOccurred())

		_, err = client.CreateRoleBinding(context.Background(), &corev1.RoleBinding{
			Id:       "test-rolebinding7",
			RoleId:   "test-role",
			Subjects: []string{"test-subject"},
		})
		Expect(err).To(HaveOccurred())
		Expect(status.Code(err)).To(Equal(codes.AlreadyExists))

		_, err = client.DeleteRoleBinding(context.Background(), &corev1.Reference{
			Id: "test-rolebinding7",
		})
		Expect(err).NotTo(HaveOccurred())
	})

	//#endregion

})
