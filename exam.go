package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	proto_mail_history "github.com/anhvanhoa/sf-proto/gen/mail_history/v1"
	proto_mail_provider "github.com/anhvanhoa/sf-proto/gen/mail_provider/v1"
	proto_mail_status "github.com/anhvanhoa/sf-proto/gen/mail_status/v1"
	proto_mail_tmpl "github.com/anhvanhoa/sf-proto/gen/mail_tmpl/v1"
	proto_status_history "github.com/anhvanhoa/sf-proto/gen/status_history/v1"
	proto_type_mail "github.com/anhvanhoa/sf-proto/gen/type_mail/v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverAddress string

func init() {
	viper.SetConfigFile("dev.config.yaml")
	viper.ReadInConfig()
	serverAddress = fmt.Sprintf("%s:%s", viper.GetString("host_grpc"), viper.GetString("port_grpc"))
}

type MailServiceClient struct {
	mailHistoryClient   proto_mail_history.MailHistoryServiceClient
	mailProviderClient  proto_mail_provider.MailProviderServiceClient
	mailTmplClient      proto_mail_tmpl.MailTmplServiceClient
	mailStatusClient    proto_mail_status.MailStatusServiceClient
	typeMailClient      proto_type_mail.TypeMailServiceClient
	statusHistoryClient proto_status_history.StatusHistoryServiceClient
	conn                *grpc.ClientConn
}

func NewMailServiceClient(address string) (*MailServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	return &MailServiceClient{
		mailHistoryClient:   proto_mail_history.NewMailHistoryServiceClient(conn),
		mailProviderClient:  proto_mail_provider.NewMailProviderServiceClient(conn),
		mailTmplClient:      proto_mail_tmpl.NewMailTmplServiceClient(conn),
		mailStatusClient:    proto_mail_status.NewMailStatusServiceClient(conn),
		typeMailClient:      proto_type_mail.NewTypeMailServiceClient(conn),
		statusHistoryClient: proto_status_history.NewStatusHistoryServiceClient(conn),
		conn:                conn,
	}, nil
}

func (c *MailServiceClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

// --- Helper để làm sạch input ---
func cleanInput(s string) string {
	return strings.ToValidUTF8(strings.TrimSpace(s), "")
}

// ================== Type Mail Service Tests ==================

func (c *MailServiceClient) TestCreateTypeMail() {
	fmt.Println("\n=== Test Create Type Mail ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter type mail name: ")
	name, _ := reader.ReadString('\n')
	name = cleanInput(name)

	fmt.Print("Enter created by: ")
	createdBy, _ := reader.ReadString('\n')
	createdBy = cleanInput(createdBy)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.typeMailClient.CreateTypeMail(ctx, &proto_type_mail.CreateTypeMailRequest{
		Name:      name,
		CreatedBy: createdBy,
		CreatedAt: time.Now().Format(time.RFC3339),
	})
	if err != nil {
		fmt.Printf("Error calling CreateTypeMail: %v\n", err)
		return
	}

	fmt.Printf("Create Type Mail result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("ID: %s\n", resp.TypeMail.Id)
	fmt.Printf("Name: %s\n", resp.TypeMail.Name)
	fmt.Printf("Created By: %s\n", resp.TypeMail.CreatedBy)
	fmt.Printf("Created At: %s\n", resp.TypeMail.CreatedAt)
}

func (c *MailServiceClient) TestGetTypeMail() {
	fmt.Println("\n=== Test Get Type Mail ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter type mail ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.typeMailClient.GetTypeMail(ctx, &proto_type_mail.GetTypeMailRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetTypeMail: %v\n", err)
		return
	}

	fmt.Printf("Get Type Mail result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("ID: %s\n", resp.TypeMail.Id)
	fmt.Printf("Name: %s\n", resp.TypeMail.Name)
	fmt.Printf("Created By: %s\n", resp.TypeMail.CreatedBy)
	fmt.Printf("Created At: %s\n", resp.TypeMail.CreatedAt)
}

func (c *MailServiceClient) TestGetAllTypeMail() {
	fmt.Println("\n=== Test Get All Type Mail ===")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.typeMailClient.GetAllTypeMail(ctx, &proto_type_mail.GetAllTypeMailRequest{})
	if err != nil {
		fmt.Printf("Error calling GetAllTypeMail: %v\n", err)
		return
	}

	fmt.Printf("Get All Type Mail result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Page: %d\n", resp.Page)
	fmt.Printf("Limit: %d\n", resp.Limit)
	fmt.Printf("Type Mails:\n")
	for i, typeMail := range resp.TypeMails {
		fmt.Printf("  [%d] ID: %s, Name: %s, Created By: %s\n", i+1, typeMail.Id, typeMail.Name, typeMail.CreatedBy)
	}
}

func (c *MailServiceClient) TestUpdateTypeMail() {
	fmt.Println("\n=== Test Update Type Mail ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter type mail ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	fmt.Print("Enter new name: ")
	name, _ := reader.ReadString('\n')
	name = cleanInput(name)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.typeMailClient.UpdateTypeMail(ctx, &proto_type_mail.UpdateTypeMailRequest{
		Id:   id,
		Name: name,
	})
	if err != nil {
		fmt.Printf("Error calling UpdateTypeMail: %v\n", err)
		return
	}

	fmt.Printf("Update Type Mail result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("ID: %s\n", resp.TypeMail.Id)
	fmt.Printf("Name: %s\n", resp.TypeMail.Name)
	fmt.Printf("Updated At: %s\n", resp.TypeMail.UpdatedAt)
}

func (c *MailServiceClient) TestDeleteTypeMail() {
	fmt.Println("\n=== Test Delete Type Mail ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter type mail ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.typeMailClient.DeleteTypeMail(ctx, &proto_type_mail.DeleteTypeMailRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling DeleteTypeMail: %v\n", err)
		return
	}

	fmt.Printf("Delete Type Mail result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
}

// ================== Mail Provider Service Tests ==================

func (c *MailServiceClient) TestCreateMailProvider() {
	fmt.Println("\n=== Test Create Mail Provider ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter email: ")
	email, _ := reader.ReadString('\n')
	email = cleanInput(email)

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = cleanInput(password)

	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = cleanInput(username)

	fmt.Print("Enter port: ")
	portStr, _ := reader.ReadString('\n')
	portStr = cleanInput(portStr)
	port := int32(587)
	if portStr != "" {
		if p, err := strconv.Atoi(portStr); err == nil {
			port = int32(p)
		}
	}

	fmt.Print("Enter host: ")
	host, _ := reader.ReadString('\n')
	host = cleanInput(host)

	fmt.Print("Enter encryption: ")
	encryption, _ := reader.ReadString('\n')
	encryption = cleanInput(encryption)

	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')
	name = cleanInput(name)

	fmt.Print("Enter type ID: ")
	typeId, _ := reader.ReadString('\n')
	typeId = cleanInput(typeId)

	fmt.Print("Enter created by: ")
	createdBy, _ := reader.ReadString('\n')
	createdBy = cleanInput(createdBy)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.mailProviderClient.CreateMailProvider(ctx, &proto_mail_provider.CreateMailProviderRequest{
		Email:      email,
		Password:   password,
		UserName:   username,
		Port:       port,
		Host:       host,
		Encryption: encryption,
		Name:       name,
		TypeId:     typeId,
		CreatedBy:  createdBy,
	})
	if err != nil {
		fmt.Printf("Error calling CreateMailProvider: %v\n", err)
		return
	}

	fmt.Printf("Create Mail Provider result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("Email: %s\n", resp.MailProvider.Email)
	fmt.Printf("Username: %s\n", resp.MailProvider.UserName)
	fmt.Printf("Host: %s\n", resp.MailProvider.Host)
	fmt.Printf("Port: %d\n", resp.MailProvider.Port)
	fmt.Printf("Name: %s\n", resp.MailProvider.Name)
	fmt.Printf("Created By: %s\n", resp.MailProvider.CreatedBy)
}

func (c *MailServiceClient) TestGetMailProvider() {
	fmt.Println("\n=== Test Get Mail Provider ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter email: ")
	email, _ := reader.ReadString('\n')
	email = cleanInput(email)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.mailProviderClient.GetMailProvider(ctx, &proto_mail_provider.GetMailProviderRequest{
		Email: email,
	})
	if err != nil {
		fmt.Printf("Error calling GetMailProvider: %v\n", err)
		return
	}

	fmt.Printf("Get Mail Provider result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("Email: %s\n", resp.MailProvider.Email)
	fmt.Printf("Username: %s\n", resp.MailProvider.UserName)
	fmt.Printf("Host: %s\n", resp.MailProvider.Host)
	fmt.Printf("Port: %d\n", resp.MailProvider.Port)
	fmt.Printf("Name: %s\n", resp.MailProvider.Name)
}

func (c *MailServiceClient) TestGetAllMailProvider() {
	fmt.Println("\n=== Test Get All Mail Provider ===")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.mailProviderClient.GetAllMailProvider(ctx, &proto_mail_provider.GetAllMailProviderRequest{})
	if err != nil {
		fmt.Printf("Error calling GetAllMailProvider: %v\n", err)
		return
	}

	fmt.Printf("Get All Mail Provider result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Page: %d\n", resp.Page)
	fmt.Printf("Limit: %d\n", resp.Limit)
	fmt.Printf("Mail Providers:\n")
	for i, provider := range resp.MailProviders {
		fmt.Printf("  [%d] Email: %s, Name: %s, Host: %s\n", i+1, provider.Email, provider.Name, provider.Host)
	}
}

func (c *MailServiceClient) TestUpdateMailProvider() {
	fmt.Println("\n=== Test Update Mail Provider ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter email: ")
	email, _ := reader.ReadString('\n')
	email = cleanInput(email)

	fmt.Print("Enter new password: ")
	password, _ := reader.ReadString('\n')
	password = cleanInput(password)

	fmt.Print("Enter new username: ")
	username, _ := reader.ReadString('\n')
	username = cleanInput(username)

	fmt.Print("Enter new port: ")
	portStr, _ := reader.ReadString('\n')
	portStr = cleanInput(portStr)
	port := int32(587)
	if portStr != "" {
		if p, err := strconv.Atoi(portStr); err == nil {
			port = int32(p)
		}
	}

	fmt.Print("Enter new host: ")
	host, _ := reader.ReadString('\n')
	host = cleanInput(host)

	fmt.Print("Enter new encryption: ")
	encryption, _ := reader.ReadString('\n')
	encryption = cleanInput(encryption)

	fmt.Print("Enter new name: ")
	name, _ := reader.ReadString('\n')
	name = cleanInput(name)

	fmt.Print("Enter new type ID: ")
	typeId, _ := reader.ReadString('\n')
	typeId = cleanInput(typeId)

	fmt.Print("Enter status (default active): ")
	status, _ := reader.ReadString('\n')
	status = cleanInput(status)
	if status == "" {
		status = "active"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.mailProviderClient.UpdateMailProvider(ctx, &proto_mail_provider.UpdateMailProviderRequest{
		Email:      email,
		Password:   password,
		UserName:   username,
		Port:       port,
		Host:       host,
		Encryption: encryption,
		Name:       name,
		TypeId:     typeId,
		Status:     status,
	})
	if err != nil {
		fmt.Printf("Error calling UpdateMailProvider: %v\n", err)
		return
	}

	fmt.Printf("Update Mail Provider result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("Email: %s\n", resp.MailProvider.Email)
	fmt.Printf("Username: %s\n", resp.MailProvider.UserName)
	fmt.Printf("Host: %s\n", resp.MailProvider.Host)
	fmt.Printf("Port: %d\n", resp.MailProvider.Port)
	fmt.Printf("Name: %s\n", resp.MailProvider.Name)
	fmt.Printf("Type ID: %s\n", resp.MailProvider.TypeId)
	fmt.Printf("Updated At: %s\n", resp.MailProvider.UpdatedAt)
}

func (c *MailServiceClient) TestDeleteMailProvider() {
	fmt.Println("\n=== Test Delete Mail Provider ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter email to delete: ")
	email, _ := reader.ReadString('\n')
	email = cleanInput(email)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.mailProviderClient.DeleteMailProvider(ctx, &proto_mail_provider.DeleteMailProviderRequest{
		Email: email,
	})
	if err != nil {
		fmt.Printf("Error calling DeleteMailProvider: %v\n", err)
		return
	}

	fmt.Printf("Delete Mail Provider result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
}

// ================== Mail Template Service Tests ==================

func (c *MailServiceClient) TestCreateMailTemplate() {
	fmt.Println("\n=== Test Create Mail Template ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter template ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	fmt.Print("Enter template name: ")
	name, _ := reader.ReadString('\n')
	name = cleanInput(name)

	fmt.Print("Enter subject: ")
	subject, _ := reader.ReadString('\n')
	subject = cleanInput(subject)

	fmt.Print("Enter keys (comma separated): ")
	keysStr, _ := reader.ReadString('\n')
	keysStr = cleanInput(keysStr)
	var keys []string
	if keysStr != "" {
		keys = strings.Split(keysStr, ",")
		for i, key := range keys {
			keys[i] = strings.TrimSpace(key)
		}
	}

	fmt.Print("Enter body: ")
	body, _ := reader.ReadString('\n')
	body = cleanInput(body)

	fmt.Print("Enter created by: ")
	createdBy, _ := reader.ReadString('\n')
	createdBy = cleanInput(createdBy)

	fmt.Print("Enter status (default active): ")
	status, _ := reader.ReadString('\n')
	status = cleanInput(status)
	if status == "" {
		status = "active"
	}

	fmt.Print("Enter provider email: ")
	providerEmail, _ := reader.ReadString('\n')
	providerEmail = cleanInput(providerEmail)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.mailTmplClient.CreateMailTmpl(ctx, &proto_mail_tmpl.CreateMailTmplRequest{
		Id:            id,
		Name:          name,
		Subject:       subject,
		Keys:          keys,
		Body:          body,
		CreatedBy:     createdBy,
		Status:        status,
		ProviderEmail: providerEmail,
	})
	if err != nil {
		fmt.Printf("Error calling CreateMailTmpl: %v\n", err)
		return
	}

	fmt.Printf("Create Mail Template result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("ID: %s\n", resp.MailTmpl.Id)
	fmt.Printf("Name: %s\n", resp.MailTmpl.Name)
	fmt.Printf("Subject: %s\n", resp.MailTmpl.Subject)
	fmt.Printf("Keys: %v\n", resp.MailTmpl.Keys)
	fmt.Printf("Body: %s\n", resp.MailTmpl.Body)
	fmt.Printf("Provider Email: %s\n", resp.MailTmpl.ProviderEmail)
	fmt.Printf("Created By: %s\n", resp.MailTmpl.CreatedBy)
	fmt.Printf("Created At: %s\n", resp.MailTmpl.CreatedAt)
}

func (c *MailServiceClient) TestGetMailTemplate() {
	fmt.Println("\n=== Test Get Mail Template ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter template ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.mailTmplClient.GetMailTmpl(ctx, &proto_mail_tmpl.GetMailTmplRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetMailTmpl: %v\n", err)
		return
	}

	fmt.Printf("Get Mail Template result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("ID: %s\n", resp.MailTmpl.Id)
	fmt.Printf("Name: %s\n", resp.MailTmpl.Name)
	fmt.Printf("Subject: %s\n", resp.MailTmpl.Subject)
	fmt.Printf("Keys: %v\n", resp.MailTmpl.Keys)
	fmt.Printf("Body: %s\n", resp.MailTmpl.Body)
	fmt.Printf("Provider Email: %s\n", resp.MailTmpl.ProviderEmail)
	fmt.Printf("Created By: %s\n", resp.MailTmpl.CreatedBy)
	fmt.Printf("Created At: %s\n", resp.MailTmpl.CreatedAt)
}

func (c *MailServiceClient) TestGetAllMailTemplate() {
	fmt.Println("\n=== Test Get All Mail Template ===")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.mailTmplClient.GetAllMailTmpl(ctx, &proto_mail_tmpl.GetAllMailTmplRequest{})
	if err != nil {
		fmt.Printf("Error calling GetAllMailTmpl: %v\n", err)
		return
	}

	fmt.Printf("Get All Mail Template result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Mail Templates:\n")
	for i, template := range resp.MailTmpls {
		fmt.Printf("  [%d] ID: %s, Name: %s, Subject: %s, Created By: %s\n", i+1, template.Id, template.Name, template.Subject, template.CreatedBy)
	}
}

func (c *MailServiceClient) TestUpdateMailTemplate() {
	fmt.Println("\n=== Test Update Mail Template ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter template ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	fmt.Print("Enter new name: ")
	name, _ := reader.ReadString('\n')
	name = cleanInput(name)

	fmt.Print("Enter new subject: ")
	subject, _ := reader.ReadString('\n')
	subject = cleanInput(subject)

	fmt.Print("Enter new keys (comma separated): ")
	keysStr, _ := reader.ReadString('\n')
	keysStr = cleanInput(keysStr)
	var keys []string
	if keysStr != "" {
		keys = strings.Split(keysStr, ",")
		for i, key := range keys {
			keys[i] = strings.TrimSpace(key)
		}
	}

	fmt.Print("Enter new body: ")
	body, _ := reader.ReadString('\n')
	body = cleanInput(body)

	fmt.Print("Enter status (default active): ")
	status, _ := reader.ReadString('\n')
	status = cleanInput(status)
	if status == "" {
		status = "active"
	}

	fmt.Print("Enter provider email: ")
	providerEmail, _ := reader.ReadString('\n')
	providerEmail = cleanInput(providerEmail)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.mailTmplClient.UpdateMailTmpl(ctx, &proto_mail_tmpl.UpdateMailTmplRequest{
		Id:            id,
		Name:          name,
		Subject:       subject,
		Keys:          keys,
		Body:          body,
		Status:        status,
		ProviderEmail: providerEmail,
	})
	if err != nil {
		fmt.Printf("Error calling UpdateMailTmpl: %v\n", err)
		return
	}

	fmt.Printf("Update Mail Template result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
}

func (c *MailServiceClient) TestDeleteMailTemplate() {
	fmt.Println("\n=== Test Delete Mail Template ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter template ID to delete: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.mailTmplClient.DeleteMailTmpl(ctx, &proto_mail_tmpl.DeleteMailTmplRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling DeleteMailTmpl: %v\n", err)
		return
	}

	fmt.Printf("Delete Mail Template result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
}

// ================== Mail History Service Tests ==================

func (c *MailServiceClient) TestCreateMailHistory() {
	fmt.Println("\n=== Test Create Mail History ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter mail history ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	fmt.Print("Enter template ID: ")
	templateId, _ := reader.ReadString('\n')
	templateId = cleanInput(templateId)

	fmt.Print("Enter subject: ")
	subject, _ := reader.ReadString('\n')
	subject = cleanInput(subject)

	fmt.Print("Enter body: ")
	body, _ := reader.ReadString('\n')
	body = cleanInput(body)

	fmt.Print("Enter recipients (comma separated): ")
	tosStr, _ := reader.ReadString('\n')
	tosStr = cleanInput(tosStr)
	var tos []string
	if tosStr != "" {
		tos = strings.Split(tosStr, ",")
		for i, to := range tos {
			tos[i] = strings.TrimSpace(to)
		}
	}

	fmt.Print("Enter data (JSON format): ")
	data, _ := reader.ReadString('\n')
	data = cleanInput(data)

	fmt.Print("Enter email provider: ")
	emailProvider, _ := reader.ReadString('\n')
	emailProvider = cleanInput(emailProvider)

	fmt.Print("Enter created by: ")
	createdBy, _ := reader.ReadString('\n')
	createdBy = cleanInput(createdBy)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.mailHistoryClient.CreateMailHistory(ctx, &proto_mail_history.CreateMailHistoryRequest{
		Id:            id,
		TemplateId:    templateId,
		Subject:       subject,
		Body:          body,
		Tos:           tos,
		Data:          data,
		EmailProvider: emailProvider,
		CreatedBy:     createdBy,
	})
	if err != nil {
		fmt.Printf("Error calling CreateMailHistory: %v\n", err)
		return
	}

	fmt.Printf("Create Mail History result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("ID: %s\n", resp.MailHistory.Id)
	fmt.Printf("Template ID: %s\n", resp.MailHistory.TemplateId)
	fmt.Printf("Subject: %s\n", resp.MailHistory.Subject)
	fmt.Printf("Recipients: %v\n", resp.MailHistory.Tos)
	fmt.Printf("Email Provider: %s\n", resp.MailHistory.EmailProvider)
	fmt.Printf("Created By: %s\n", resp.MailHistory.CreatedBy)
	fmt.Printf("Created At: %s\n", resp.MailHistory.CreatedAt)
}

func (c *MailServiceClient) TestGetMailHistory() {
	fmt.Println("\n=== Test Get Mail History ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter mail history ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.mailHistoryClient.GetMailHistory(ctx, &proto_mail_history.GetMailHistoryRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetMailHistory: %v\n", err)
		return
	}

	fmt.Printf("Get Mail History result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("ID: %s\n", resp.MailHistory.Id)
	fmt.Printf("Template ID: %s\n", resp.MailHistory.TemplateId)
	fmt.Printf("Subject: %s\n", resp.MailHistory.Subject)
	fmt.Printf("Body: %s\n", resp.MailHistory.Body)
	fmt.Printf("Recipients: %v\n", resp.MailHistory.Tos)
	fmt.Printf("Data: %s\n", resp.MailHistory.Data)
	fmt.Printf("Email Provider: %s\n", resp.MailHistory.EmailProvider)
	fmt.Printf("Created By: %s\n", resp.MailHistory.CreatedBy)
	fmt.Printf("Created At: %s\n", resp.MailHistory.CreatedAt)
	fmt.Printf("Updated At: %s\n", resp.MailHistory.UpdatedAt)
}

func (c *MailServiceClient) TestGetAllMailHistory() {
	fmt.Println("\n=== Test Get All Mail History ===")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.mailHistoryClient.GetAllMailHistory(ctx, &proto_mail_history.GetAllMailHistoryRequest{})
	if err != nil {
		fmt.Printf("Error calling GetAllMailHistory: %v\n", err)
		return
	}

	fmt.Printf("Get All Mail History result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Page: %d\n", resp.Page)
	fmt.Printf("Limit: %d\n", resp.Limit)
	fmt.Printf("Mail Histories:\n")
	for i, history := range resp.MailHistories {
		fmt.Printf("  [%d] ID: %s, Subject: %s, Email Provider: %s, Created By: %s\n", i+1, history.Id, history.Subject, history.EmailProvider, history.CreatedBy)
	}
}

func (c *MailServiceClient) TestUpdateMailHistory() {
	fmt.Println("\n=== Test Update Mail History ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter mail history ID to update: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	fmt.Print("Enter template ID: ")
	templateId, _ := reader.ReadString('\n')
	templateId = cleanInput(templateId)

	fmt.Print("Enter new subject: ")
	subject, _ := reader.ReadString('\n')
	subject = cleanInput(subject)

	fmt.Print("Enter new body: ")
	body, _ := reader.ReadString('\n')
	body = cleanInput(body)

	fmt.Print("Enter recipients (comma separated): ")
	tosStr, _ := reader.ReadString('\n')
	tosStr = cleanInput(tosStr)
	var tos []string
	if tosStr != "" {
		tos = strings.Split(tosStr, ",")
		for i, to := range tos {
			tos[i] = strings.TrimSpace(to)
		}
	}

	fmt.Print("Enter data (JSON format): ")
	data, _ := reader.ReadString('\n')
	data = cleanInput(data)

	fmt.Print("Enter email provider: ")
	emailProvider, _ := reader.ReadString('\n')
	emailProvider = cleanInput(emailProvider)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.mailHistoryClient.UpdateMailHistory(ctx, &proto_mail_history.UpdateMailHistoryRequest{
		Id:            id,
		TemplateId:    templateId,
		Subject:       subject,
		Body:          body,
		Tos:           tos,
		Data:          data,
		EmailProvider: emailProvider,
	})
	if err != nil {
		fmt.Printf("Error calling UpdateMailHistory: %v\n", err)
		return
	}

	fmt.Printf("Update Mail History result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("ID: %s\n", resp.MailHistory.Id)
	fmt.Printf("Template ID: %s\n", resp.MailHistory.TemplateId)
	fmt.Printf("Subject: %s\n", resp.MailHistory.Subject)
	fmt.Printf("Recipients: %v\n", resp.MailHistory.Tos)
	fmt.Printf("Email Provider: %s\n", resp.MailHistory.EmailProvider)
	fmt.Printf("Updated At: %s\n", resp.MailHistory.UpdatedAt)
}

func (c *MailServiceClient) TestDeleteMailHistory() {
	fmt.Println("\n=== Test Delete Mail History ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter mail history ID to delete: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.mailHistoryClient.DeleteMailHistory(ctx, &proto_mail_history.DeleteMailHistoryRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling DeleteMailHistory: %v\n", err)
		return
	}

	fmt.Printf("Delete Mail History result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
}

// ================== Mail Status Service Tests ==================

func (c *MailServiceClient) TestGetMailStatus() {
	fmt.Println("\n=== Test Get Mail Status ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter status: ")
	status, _ := reader.ReadString('\n')
	status = cleanInput(status)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.mailStatusClient.GetMailStatus(ctx, &proto_mail_status.GetMailStatusRequest{
		Status: status,
	})
	if err != nil {
		fmt.Printf("Error calling GetMailStatus: %v\n", err)
		return
	}

	fmt.Printf("Get Mail Status result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("Status: %s\n", resp.MailStatus.Status)
	fmt.Printf("Name: %s\n", resp.MailStatus.Name)
	fmt.Printf("Created At: %s\n", resp.MailStatus.CreatedAt)
}

func (c *MailServiceClient) TestGetAllMailStatus() {
	fmt.Println("\n=== Test Get All Mail Status ===")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.mailStatusClient.GetAllMailStatus(ctx, &proto_mail_status.GetAllMailStatusRequest{})
	if err != nil {
		fmt.Printf("Error calling GetAllMailStatus: %v\n", err)
		return
	}

	fmt.Printf("Get All Mail Status result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("Total: %d\n", resp.Total)
	fmt.Printf("Page: %d\n", resp.Page)
	fmt.Printf("Limit: %d\n", resp.Limit)
	fmt.Printf("Mail Statuses:\n")
	for i, status := range resp.MailStatuses {
		fmt.Printf("  [%d] Status: %s, Name: %s, Created At: %s\n", i+1, status.Status, status.Name, status.CreatedAt)
	}
}

// ================== Status History Service Tests ==================

func (c *MailServiceClient) TestCreateStatusHistory() {
	fmt.Println("\n=== Test Create Status History ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter mail history ID: ")
	mailHistoryId, _ := reader.ReadString('\n')
	mailHistoryId = cleanInput(mailHistoryId)

	fmt.Print("Enter status: ")
	status, _ := reader.ReadString('\n')
	status = cleanInput(status)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.statusHistoryClient.CreateStatusHistory(ctx, &proto_status_history.CreateStatusHistoryRequest{
		MailHistoryId: mailHistoryId,
		Status:        status,
		Message:       "Status updated",
	})
	if err != nil {
		fmt.Printf("Error calling CreateStatusHistory: %v\n", err)
		return
	}

	fmt.Printf("Create Status History result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("Mail History ID: %s\n", resp.StatusHistory.MailHistoryId)
	fmt.Printf("Status: %s\n", resp.StatusHistory.Status)
	fmt.Printf("Message: %s\n", resp.StatusHistory.Message)
	fmt.Printf("Created At: %s\n", resp.StatusHistory.CreatedAt)
}

func (c *MailServiceClient) TestGetStatusHistoryByMailHistoryId() {
	fmt.Println("\n=== Test Get Status History By Mail History ID ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter mail history ID: ")
	mailHistoryId, _ := reader.ReadString('\n')
	mailHistoryId = cleanInput(mailHistoryId)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.statusHistoryClient.GetStatusHistoryByMailHistoryId(ctx, &proto_status_history.GetStatusHistoryByMailHistoryIdRequest{
		MailHistoryId: mailHistoryId,
	})
	if err != nil {
		fmt.Printf("Error calling GetStatusHistoryByMailHistoryId: %v\n", err)
		return
	}

	fmt.Printf("Get Status History By Mail History ID result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("Status Histories:\n")
	for i, history := range resp.StatusHistories {
		fmt.Printf("  [%d] Status: %s, Message: %s, Created At: %s\n", i+1, history.Status, history.Message, history.CreatedAt)
	}
}

func printMainMenu() {
	fmt.Println("\n=== gRPC Mail Service Test Client ===")
	fmt.Println("1. Type Mail Service")
	fmt.Println("2. Mail Provider Service")
	fmt.Println("3. Mail Template Service")
	fmt.Println("4. Mail History Service")
	fmt.Println("5. Mail Status Service")
	fmt.Println("6. Status History Service")
	fmt.Println("0. Exit")
	fmt.Print("Enter your choice: ")
}

func printTypeMailMenu() {
	fmt.Println("\n=== Type Mail Service ===")
	fmt.Println("1. Create Type Mail")
	fmt.Println("2. Get Type Mail")
	fmt.Println("3. Get All Type Mail")
	fmt.Println("4. Update Type Mail")
	fmt.Println("5. Delete Type Mail")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func printMailProviderMenu() {
	fmt.Println("\n=== Mail Provider Service ===")
	fmt.Println("1. Create Mail Provider")
	fmt.Println("2. Get Mail Provider")
	fmt.Println("3. Get All Mail Provider")
	fmt.Println("4. Update Mail Provider")
	fmt.Println("5. Delete Mail Provider")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func printMailTemplateMenu() {
	fmt.Println("\n=== Mail Template Service ===")
	fmt.Println("1. Create Mail Template")
	fmt.Println("2. Get Mail Template")
	fmt.Println("3. Get All Mail Template")
	fmt.Println("4. Update Mail Template")
	fmt.Println("5. Delete Mail Template")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func printMailHistoryMenu() {
	fmt.Println("\n=== Mail History Service ===")
	fmt.Println("1. Create Mail History")
	fmt.Println("2. Get Mail History")
	fmt.Println("3. Get All Mail History")
	fmt.Println("4. Update Mail History")
	fmt.Println("5. Delete Mail History")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func printMailStatusMenu() {
	fmt.Println("\n=== Mail Status Service ===")
	fmt.Println("1. Get Mail Status")
	fmt.Println("2. Get All Mail Status")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func printStatusHistoryMenu() {
	fmt.Println("\n=== Status History Service ===")
	fmt.Println("1. Create Status History")
	fmt.Println("2. Get Status History By Mail History ID")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func main() {
	address := serverAddress
	if len(os.Args) > 1 {
		address = os.Args[1]
	}

	fmt.Printf("Connecting to gRPC server at %s...\n", address)
	client, err := NewMailServiceClient(address)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer client.Close()

	fmt.Println("Connected successfully!")

	reader := bufio.NewReader(os.Stdin)

	for {
		printMainMenu()
		choice, _ := reader.ReadString('\n')
		choice = cleanInput(choice)

		switch choice {
		case "1":
			for {
				printTypeMailMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateTypeMail()
				case "2":
					client.TestGetTypeMail()
				case "3":
					client.TestGetAllTypeMail()
				case "4":
					client.TestUpdateTypeMail()
				case "5":
					client.TestDeleteTypeMail()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "2":
			for {
				printMailProviderMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateMailProvider()
				case "2":
					client.TestGetMailProvider()
				case "3":
					client.TestGetAllMailProvider()
				case "4":
					client.TestUpdateMailProvider()
				case "5":
					client.TestDeleteMailProvider()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "3":
			// Mail Template Service
			for {
				printMailTemplateMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateMailTemplate()
				case "2":
					client.TestGetMailTemplate()
				case "3":
					client.TestGetAllMailTemplate()
				case "4":
					client.TestUpdateMailTemplate()
				case "5":
					client.TestDeleteMailTemplate()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "4":
			// Mail History Service
			for {
				printMailHistoryMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateMailHistory()
				case "2":
					client.TestGetMailHistory()
				case "3":
					client.TestGetAllMailHistory()
				case "4":
					client.TestUpdateMailHistory()
				case "5":
					client.TestDeleteMailHistory()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "5":
			// Mail Status Service
			for {
				printMailStatusMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestGetMailStatus()
				case "2":
					client.TestGetAllMailStatus()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "6":
			// Status History Service
			for {
				printStatusHistoryMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateStatusHistory()
				case "2":
					client.TestGetStatusHistoryByMailHistoryId()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "0":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
