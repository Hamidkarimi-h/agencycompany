package utils

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type InputHelper struct {
    reader *bufio.Reader
}

func NewInputHelper() *InputHelper {
    return &InputHelper{reader: bufio.NewReader(os.Stdin)}
}

func (h *InputHelper) ReadString(prompt string) (string, error) {
    fmt.Print(prompt)
    input, err := h.reader.ReadString('\n')
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(input), nil
}

func (h *InputHelper) ReadRequired(prompt string) (string, error) {
    for {
        input, err := h.ReadString(prompt)
        if err != nil {
            return "", err
        }
        if input == "" {
            fmt.Println("⚠️  This field is required.")
            continue
        }
        return input, nil
    }
}

func (h *InputHelper) ReadOptional(prompt string) (string, error) {
    return h.ReadString(prompt)
}

func (h *InputHelper) ReadWithDefault(prompt, defaultValue string) (string, error) {
    displayPrompt := prompt
    if defaultValue != "" {
        displayPrompt = fmt.Sprintf("%s [%s]: ", strings.TrimSuffix(prompt, ": "), defaultValue)
    } else {
        displayPrompt = fmt.Sprintf("%s (press Enter to skip): ", strings.TrimSuffix(prompt, ": "))
    }

    input, err := h.ReadString(displayPrompt)
    if err != nil {
        return "", err
    }

    if input == "" {
        return defaultValue, nil
    }
    return input, nil
}

func (h *InputHelper) ReadUint32(prompt string) (*uint32, error) {
    input, err := h.ReadString(prompt)
    if err != nil {
        return nil, err
    }

    if input == "" {
        return nil, nil
    }

    count, err := strconv.ParseUint(input, 10, 32)
    if err != nil {
        return nil, err
    }

    val := uint32(count)
    return &val, nil
}

func (h *InputHelper) ReadUint(prompt string) (uint, error) {
    input, err := h.ReadString(prompt)
    if err != nil {
        return 0, err
    }

    u64, err := strconv.ParseUint(input, 10, 0)
    if err != nil {
        return 0, err
    }

    return uint(u64), nil
}